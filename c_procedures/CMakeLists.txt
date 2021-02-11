cmake_minimum_required(VERSION 3.16)
project(procedures C)

set(CMAKE_C_FLAGS "${CMAKE_C_FLAGS} -O2")
set(CMAKE_POSITION_INDEPENDENT_CODE ON)
set(CMAKE_C_STANDARD 11)

include_directories(/usr/local/include)
link_directories(/usr/local/lib)

add_library(procedures SHARED procedures.c)
target_link_libraries(procedures libmsgpuck.a uuid)
