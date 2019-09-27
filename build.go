package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/build"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

const (
	// ProtobufVersion is the version of the protocol buffers.
	ProtobufVersion = "3.7.1"
	// PrototoolVersion specifies the version of the prototool.
	PrototoolVersion = "1.6.0"
)

var (
	// ProtoOS specifies the operative system for prototool.
	ProtoOS = getProtoOS()
	// ProtocFilename is the name of the protoc binary.
	ProtocFilename = fmt.Sprintf("protoc-%s-%s-x86_64.zip", ProtobufVersion, ProtoOS)
)

var (
	gitMail = flag.String("git-author-mail", "cs3org-bot@hugo.labkode.com", "Git author mail")
	gitName = flag.String("git-author-name", "cs3org-bot", "Git author name")
	gitSSH  = flag.Bool("git-ssh", false, "Use git protocol instead of https for cloning repos")

	_all = flag.Bool("all", false, "Compile, build and publish for all available languages, mean to be run in CI platform like Drone")

	_buildProto = flag.Bool("build-proto", false, "Compile Protobuf definitions")

	_buildGo = flag.Bool("build-go", false, "Build Go library")
	_pushGo  = flag.Bool("push-go", false, "Push Go library to github.com/cs3org/go-cs3apis")

	_buildPython = flag.Bool("build-python", false, "Build Python library")
	_pushPython  = flag.Bool("push-python", false, "Push Python library to github.com/cs3org/python-cs3apis")

	_buildJs = flag.Bool("build-js", false, "Build Js library")
	_pushJs  = flag.Bool("push-js", false, "Push Js library to github.com/cs3org/js-cs3apis")
)

func init() {
	flag.Parse()

	if *_all {
		*_buildProto = true
		*_buildGo = true
		*_buildPython = true
		*_buildJs = true

		*_pushGo = true
		*_pushPython = true
		*_pushJs = true
	}
}

func getProtoOS() string {
	switch runtime.GOOS {
	case "darwin":
		return "osx"
	case "linux":
		return "linux"
	default:
		panic("no build procedure for " + runtime.GOOS)
	}
}

func clone(repo, dir string) {
	repo = getRepo(repo) // get git or https repo location
	cmd := exec.Command("git", "clone", "--quiet", repo)
	cmd.Dir = dir
	run(cmd)
}

func checkout(branch, dir string) {
	// See https://stackoverflow.com/questions/26961371/switch-on-another-branch-create-if-not-exists-without-checking-if-already-exi
	cmd := exec.Command("bash", "-c", fmt.Sprintf("git checkout %s || git checkout -b %s", branch, branch))
	cmd.Dir = dir
	run(cmd)
}

func update(dir string) error {
	cmd := exec.Command("git", "pull", "--quiet")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = dir
	return cmd.Run()
}

func isRepoDirty(repo string) bool {
	cmd := exec.Command("git", "status", "-s")
	cmd.Dir = repo
	changes := runAndGet(cmd)
	if changes != "" {
		fmt.Println("repo is dirty")
		fmt.Println(changes)
	}
	return changes != ""
}

func getCommitID(dir string) string {
	if os.Getenv("BUILD_GIT_COMMIT") != "" {
		return os.Getenv("BUILD_GIT_COMMIT")
	}

	cmd := exec.Command("git", "rev-parse", "HEAD")
	cmd.Dir = dir
	commit := runAndGet(cmd)
	return commit
}

func getRepo(repo string) string {
	if *gitSSH {
		return fmt.Sprintf("git@github.com:%s", repo)
	}
	return fmt.Sprintf("https://github.com/%s", repo)
}

func commit(repo, msg string) {
	// set correct author name and mail
	cmd := exec.Command("git", "config", "user.email", *gitMail)
	cmd.Dir = repo
	run(cmd)

	cmd = exec.Command("git", "config", "user.name", *gitName)
	cmd.Dir = repo
	run(cmd)

	// check if repo is dirty
	if !isRepoDirty(repo) {
		// nothing to do
		return
	}

	cmd = exec.Command("git", "add", ".")
	cmd.Dir = repo
	run(cmd)

	cmd = exec.Command("git", "commit", "-m", msg)
	cmd.Dir = repo
	run(cmd)
}

func push(repo string) {
	protoBranch := getGitBranch(".")
	cmd := exec.Command("git", "push", "--set-upstream", "origin", protoBranch)
	cmd.Dir = repo
	run(cmd)
}

func getGitBranch(repo string) string {
	// check if branch is provided by env variable
	if os.Getenv("BUILD_GIT_BRANCH") != "" {
		return os.Getenv("BUILD_GIT_BRANCH")
	}

	// obtain branch from repo
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	cmd.Dir = repo
	branch := runAndGet(cmd)
	return branch
}

// getVersionFromGit returns a version string that identifies the currently
// checked out git commit.
func getVersionFromGit(repodir string) string {
	cmd := exec.Command("git", "describe",
		"--long", "--tags", "--dirty", "--always")
	cmd.Dir = repodir
	out, err := cmd.Output()
	if err != nil {
		panic(fmt.Sprintf("git describe returned error: %v\n", err))
	}

	version := strings.TrimSpace(string(out))
	return version
}

func run(cmd *exec.Cmd) {
	var b bytes.Buffer
	mw := io.MultiWriter(os.Stdout, &b)
	cmd.Stdout = mw
	cmd.Stderr = mw
	err := cmd.Run()
	fmt.Println(cmd.Dir, cmd.Args)
	fmt.Println(b.String())
	if err != nil {
		fmt.Println("ERROR: ", err.Error())
		os.Exit(1)
	}
}

func runAndGet(cmd *exec.Cmd) string {
	var b bytes.Buffer
	mw := io.MultiWriter(os.Stdout, &b)
	cmd.Stderr = mw
	out, err := cmd.Output()
	fmt.Println(cmd.Dir, cmd.Args)
	fmt.Println(b.String())
	if err != nil {
		fmt.Println("ERROR: ", err.Error())
		os.Exit(1)
	}
	return strings.TrimSpace(string(out))
}

// Works with Go 1.8+
// https://stackoverflow.com/questions/32649770/how-to-get-current-gopath-from-code
func getGoPath() string {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = build.Default.GOPATH
	}
	return gopath
}

func sed(dir, suffix, old, new string) {
	err := filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		if strings.HasSuffix(path, suffix) {
			data, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			newData := strings.ReplaceAll(string(data), old, new)
			err = ioutil.WriteFile(path, []byte(newData), 0)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func find(patterns ...string) []string {
	var files []string
	for _, p := range patterns {
		fs, err := filepath.Glob(p)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		files = append(files, fs...)
	}
	return files
}

func buildProto() {
	dir := "."
	cmd := exec.Command("prototool", "compile")
	cmd.Dir = dir
	run(cmd)

	cmd = exec.Command("protolock", "status")
	cmd.Dir = dir
	run(cmd)

	// lint
	cmd = exec.Command("prototool", "format", "-w")
	cmd.Dir = dir
	run(cmd)
	cmd = exec.Command("prototool", "lint")
	cmd.Dir = dir
	run(cmd)
	cmd = exec.Command("go", "run", "tools/check-license/check-license.go")
	cmd.Dir = dir
	run(cmd)

	os.RemoveAll("docs")
	os.MkdirAll("docs", 0755)

	files := find("cs3/*/*.proto", "cs3/*/*/*.proto")
	fmt.Println(files)

	args := []string{"--doc_out=./docs", "--doc_opt=html,index.html"}
	args = append(args, files...)
	cmd = exec.Command("protoc", args...)
	run(cmd)
}

func buildGo() {

	// Remove build dir
	os.RemoveAll("build/go-cs3apis")
	os.MkdirAll("build", 0755)

	// Clone Go repo and set branch to current branch
	clone("cs3org/go-cs3apis", "build")
	protoBranch := getGitBranch(".")
	goBranch := getGitBranch("build/go-cs3apis")
	fmt.Printf("Proto branch: %s\nGo branch: %s\n", protoBranch, goBranch)

	if goBranch != protoBranch {
		checkout(protoBranch, "build/go-cs3apis")
	}

	cmd := exec.Command("prototool", "generate")
	run(cmd)

	sed("build/go-cs3apis", ".go", "github.com/cs3org/go-cs3apis/build/go-cs3apis/cs3/", "github.com/cs3org/go-cs3apis/cs3/")

	if !isRepoDirty("build/go-cs3apis") {
		fmt.Println("Repo is clean, nothing to do")
	}

	// get proto repo commit id
	hash := getCommitID(".")
	repo := "build/go-cs3apis"
	msg := "Synced to https://github.com/cs3org/cs3apis/tree/" + hash
	commit(repo, msg)
}

func buildPython() {

	// Remove build dir
	os.RemoveAll("build/python-cs3apis")
	os.MkdirAll("build", 0755)

	// Clone Go repo and set branch to current branch
	clone("cs3org/python-cs3apis", "build")
	protoBranch := getGitBranch(".")
	buildBranch := getGitBranch("build/python-cs3apis")
	fmt.Printf("Proto branch: %s\nBuild branch: %s\n", protoBranch, buildBranch)

	if buildBranch != protoBranch {
		checkout(protoBranch, "build/python-cs3apis")
	}

	files := find("cs3/*/*.proto", "cs3/*/*/*.proto")

	args := []string{"-m", "grpc_tools.protoc", "--python_out=./build/python-cs3apis", "-I.", "--grpc_python_out=./build/python-cs3apis"}
	args = append(args, files...)
	cmd := exec.Command("python", args...)
	run(cmd)

	// get proto repo commit id
	hash := getCommitID(".")
	repo := "build/python-cs3apis"
	msg := "Synced to https://github.com/cs3org/cs3apis/tree/" + hash
	commit(repo, msg)
}

func buildJS() {
	// Remove build dir
	os.RemoveAll("build/js-cs3apis")
	os.MkdirAll("build", 0755)

	// Clone repo and set branch to current branch
	clone("cs3org/js-cs3apis", "build")
	protoBranch := getGitBranch(".")
	buildBranch := getGitBranch("build/js-cs3apis")
	fmt.Printf("Proto branch: %s\nBuild branch: %s\n", protoBranch, buildBranch)

	if buildBranch != protoBranch {
		checkout(protoBranch, "build/js-cs3apis")
	}

	files := find("cs3/*/*.proto", "cs3/*/*/*.proto")

	args := []string{"--js_out=import_style=commonjs:./build/js-cs3apis", "--grpc-web_out=import_style=commonjs,mode=grpcwebtext:./build/js-cs3apis/", "-I."}
	args = append(args, files...)
	cmd := exec.Command("protoc", args...)
	run(cmd)

	// get proto repo commit id
	hash := getCommitID(".")
	repo := "build/js-cs3apis"
	msg := "Synced to https://github.com/cs3org/cs3apis/tree/" + hash
	commit(repo, msg)
}

func pushPython() {
	push("build/python-cs3apis")
}

func pushGo() {
	push("build/go-cs3apis")
}

func pushJS() {
	push("build/js-cs3apis")
}

func main() {
	if *_buildProto {
		fmt.Println("Compiling and liniting protobufs ...")
		buildProto()
	}

	if *_buildGo {
		fmt.Println("Building Go ...")
		buildGo()
	}

	if *_pushGo {
		fmt.Println("Pushing Go ...")
		pushGo()
	}

	if *_buildPython {
		fmt.Println("Building Python ...")
		buildPython()
	}

	if *_pushPython {
		fmt.Println("Pushing Python ...")
		pushPython()
	}

	if *_buildJs {
		fmt.Println("Building JS ...")
		buildJS()
	}

	if *_pushJs {
		fmt.Println("Pushing Js ...")
		pushJS()
	}
}
