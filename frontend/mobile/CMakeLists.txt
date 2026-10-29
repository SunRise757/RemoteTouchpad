
qt_add_executable(frontend
    main.cpp
    mainwindow.cpp
    mainwindow.h
    mainwindow.ui

    components/Scrollbar.cpp
    components/Scrollbar.h
    components/Scrollbar.ui
)
find_package(Qt6 REQUIRED COMPONENTS WebSockets)
target_link_libraries(frontend PRIVATE Qt${QT_VERSION_MAJOR}::WebSockets)

target_include_directories(frontend PRIVATE ${CMAKE_CURRENT_SOURCE_DIR})
target_include_directories(frontend PRIVATE ${CMAKE_CURRENT_SOURCE_DIR}/components)
