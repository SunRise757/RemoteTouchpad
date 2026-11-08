cmake_minimum_required(VERSION 3.21)
project(frontend LANGUAGES CXX)

set(CMAKE_CXX_STANDARD 20)
set(CMAKE_CXX_STANDARD_REQUIRED ON)
set(CMAKE_CXX_EXTENSIONS OFF)

set(CMAKE_EXPORT_COMPILE_COMMANDS ON)

set(CMAKE_AUTOMOC ON)
set(CMAKE_AUTORCC ON)
set(CMAKE_AUTOUIC ON)

option(BUILD_MOBILE "Build the Qt/WebAssembly client" ON)
option(BUILD_PC "Build the native Linux app" OFF)

add_subdirectory(shared)

if(BUILD_MOBILE)
    add_subdirectory(mobile)
endif()

if(BUILD_PC)
    add_subdirectory(pc)
endif()