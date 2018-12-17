all:
	cp prototool-template.yaml prototool.yaml
	sed -i "s|<GOPATH>|${GOPATH}|g" prototool.yaml
	prototool compile

