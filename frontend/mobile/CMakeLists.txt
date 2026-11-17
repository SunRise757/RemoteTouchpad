find_package(Qt6 REQUIRED COMPONENTS
    Core
    Gui
    WebSockets
)

qt_standard_project_setup()

qt_add_executable(mobile
    main.cpp
    mainwindow.cpp
    mainwindow.h
    mainwindow.ui

    components/Scrollbar.cpp
    components/Scrollbar.h
    components/Scrollbar.ui
    
    views/touchpad/Touchpad.cpp
    views/touchpad/Touchpad.h
    views/touchpad/Touchpad.ui
    
    views/login/LoginForm.cpp
    views/login/LoginForm.h
    views/login/LoginForm.ui

    serverconnection.h
    serverconnection.cpp
)

target_link_libraries(mobile
    PRIVATE
        Qt::Core
        Qt::Widgets
        Qt::WebSockets
        #shared
)

target_include_directories(mobile PRIVATE
    "${CMAKE_CURRENT_SOURCE_DIR}"
    "${CMAKE_CURRENT_SOURCE_DIR}/components"
    "${CMAKE_CURRENT_SOURCE_DIR}/views"
    "${CMAKE_CURRENT_SOURCE_DIR}/views/login"
    "${CMAKE_CURRENT_SOURCE_DIR}/views/touchpad"
)

install(TARGETS mobile
    BUNDLE  DESTINATION .
    RUNTIME DESTINATION ${CMAKE_INSTALL_BINDIR}
    LIBRARY DESTINATION ${CMAKE_INSTALL_LIBDIR}
)

qt_generate_deploy_app_script(
    TARGET mobile
    OUTPUT_SCRIPT deploy_script
    NO_UNSUPPORTED_PLATFORM_ERROR
)
install(SCRIPT ${deploy_script})

add_custom_command(TARGET mobile POST_BUILD
    COMMAND ${CMAKE_COMMAND} -E copy_if_different
            "${CMAKE_CURRENT_SOURCE_DIR}/index.html"
            "$<TARGET_FILE_DIR:mobile>/index.html"
    COMMENT "Copying custom WebAssembly HTML"
)