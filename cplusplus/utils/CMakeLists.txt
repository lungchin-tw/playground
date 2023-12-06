
cmake_minimum_required(VERSION 3.8.1)
project(utils LANGUAGES CXX)

SET( CMAKE_CXX_FLAGS "-std=c++14 -Wall" )
SET( CMAKE_CXX_STANDARD 14 )


message( "========== ${PROJECT_NAME} Begin ==========")
message( "[${PROJECT_NAME}]::CMAKE_CURRENT_SOURCE_DIR: ${CMAKE_CURRENT_SOURCE_DIR}")
message( "[${PROJECT_NAME}]::PROJECT_BINARY_DIR: ${PROJECT_BINARY_DIR}")
message( "[${PROJECT_NAME}]::PROJECT_SOURCE_DIR: ${PROJECT_SOURCE_DIR}")

file( GLOB PUBLIC_HEADER ${PROJECT_SOURCE_DIR}/public/*.h )
aux_source_directory( ${PROJECT_SOURCE_DIR}/private PRIVATE_SOURCE )

add_library( ${PROJECT_NAME} STATIC ${PRIVATE_SOURCE} ${PUBLIC_HEADER})

target_include_directories(
        ${PROJECT_NAME}
        PUBLIC
        ${PROJECT_SOURCE_DIR}/public
        ${PROJECT_SOURCE_DIR}/private )

install(TARGETS ${PROJECT_NAME} ARCHIVE DESTINATION lib-static)
install(DIRECTORY public/ DESTINATION include/${PROJECT_NAME} FILES_MATCHING PATTERN "*.h")

message( "========== ${PROJECT_NAME} End ==========")
