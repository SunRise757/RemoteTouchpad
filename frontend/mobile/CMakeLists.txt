qt_add_executable(frontend
    main.cpp
    mainwindow.cpp
    mainwindow.h
    mainwindow.ui
)
find_package(Qt6 REQUIRED COMPONENTS WebSockets)
target_link_libraries(frontend PRIVATE Qt${QT_VERSION_MAJOR}::WebSockets)