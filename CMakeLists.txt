cmake_minimum_required(VERSION 3.1)

project(cs3apis LANGUAGES CXX)

find_package(protobuf REQUIRED CONFIG)
find_package(gRpc REQUIRED CONFIG)

set(CS3_BINARY_DIR ${CMAKE_CURRENT_BINARY_DIR}/cs3apis)
file(GLOB_RECURSE PROTO_FILES ${PROJECT_SOURCE_DIR}/*.proto)

set(CS3_SRC)
foreach(_f ${PROTO_FILES})
    get_filename_component(_path ${_f} DIRECTORY)
    file(RELATIVE_PATH _rel_path ${PROJECT_SOURCE_DIR} ${_path})
    get_filename_component(_name ${_f} NAME_WE)
    set(_base_name ${CS3_BINARY_DIR}/${_rel_path}/${_name})
    list(APPEND CS3_SRC ${_base_name}.pb.cc ${_base_name}.pb.h
                        ${_base_name}.grpc.pb.cc ${_base_name}.grpc.pb.h)

endforeach()

add_custom_command(
    OUTPUT ${CS3_SRC}
    COMMAND ${CMAKE_COMMAND} -E make_directory ${CS3_BINARY_DIR}
    COMMAND $<TARGET_FILE:protobuf::protoc> ARGS --proto_path=${PROJECT_SOURCE_DIR} --cpp_out=${CS3_BINARY_DIR} ${PROTO_FILES}
    COMMAND $<TARGET_FILE:protobuf::protoc> ARGS --proto_path=${PROJECT_SOURCE_DIR} --grpc_out=${CS3_BINARY_DIR} --plugin=protoc-gen-grpc=$<TARGET_FILE:gRPC::grpc_cpp_plugin> ${PROTO_FILES}
    VERBATIM
    COMMENT "Generate potobuf and grpc sources")

set_property(SOURCE ${CS3_SRC} PROPERTY SKIP_AUTOMOC ON)
add_library(cs3apis STATIC ${CS3_SRC})
target_include_directories(cs3apis PUBLIC $<BUILD_INTERFACE:${CMAKE_CURRENT_BINARY_DIR}/cs3apis/>)
target_link_libraries(cs3apis PUBLIC gRPC::grpc protobuf::libprotobuf)
target_compile_definitions(cs3apis PUBLIC _WIN32_WINNT=0x600)
