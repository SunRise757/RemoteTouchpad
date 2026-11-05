cmake_minimum_required(VERSION 3.19)
project(frontend LANGUAGES CXX)

find_package(Qt6 6.5 REQUIRED COMPONENTS Core Widgets WebSockets)

qt_standard_project_setup()

qt_add_executable(frontend
    mobile/main.cpp
    mobile/mainwindow.cpp
    mobile/mainwindow.h
    mobile/mainwindow.ui

    mobile/components/Scrollbar.cpp
    mobile/components/Scrollbar.h
    mobile/components/Scrollbar.ui
)

target_include_directories(frontend PRIVATE
    "${CMAKE_CURRENT_SOURCE_DIR}/mobile"
    "${CMAKE_CURRENT_SOURCE_DIR}/mobile/components"
)

target_link_libraries(frontend
    PRIVATE
        Qt::Core
        Qt::Widgets
        Qt::WebSockets
)

install(TARGETS frontend
    BUNDLE  DESTINATION .
    RUNTIME DESTINATION ${CMAKE_INSTALL_BINDIR}
    LIBRARY DESTINATION ${CMAKE_INSTALL_LIBDIR}
)

qt_generate_deploy_app_script(
    TARGET frontend
    OUTPUT_SCRIPT deploy_script
    NO_UNSUPPORTED_PLATFORM_ERROR
)
install(SCRIPT ${deploy_script})
