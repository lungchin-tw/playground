cmake_minimum_required(VERSION 3.8.1 FATAL_ERROR)

project( std_ptr LANGUAGES CXX)
SET( CMAKE_CXX_FLAGS "-std=c++14 -Wall" )
SET( CMAKE_CXX_STANDARD 14 )

message( "========== ${PROJECT_NAME} Begin ==========")
message( "[${PROJECT_NAME}]::CMAKE_CURRENT_SOURCE_DIR: ${CMAKE_CURRENT_SOURCE_DIR}")
message( "[${PROJECT_NAME}]::PROJECT_BINARY_DIR: ${PROJECT_BINARY_DIR}")
message( "[${PROJECT_NAME}]::PROJECT_SOURCE_DIR: ${PROJECT_SOURCE_DIR}")

add_executable( ${PROJECT_NAME} std_ptr_main.cpp )
target_link_libraries( ${PROJECT_NAME} utils )
add_subdirectory( utils )

install( TARGETS ${PROJECT_NAME} RUNTIME DESTINATION bin )

message( "========== ${PROJECT_NAME} End ==========")