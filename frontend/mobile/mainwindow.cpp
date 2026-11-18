#include "mainwindow.h"
#include "ui_mainwindow.h"

#include <QMouseEvent>

MainWindow::MainWindow(QWidget *parent)
    : QMainWindow(parent)
    , ui(new Ui::MainWindow)
{
    ui->setupUi(this);
    touchpad = ui->touchpad;
    login = ui->login;
    

    serv = new ServerConnection(this);
    if (serv->registerToken.isEmpty())
        ui->stackedWidget->setCurrentWidget(ui->login);
    else
        ui->stackedWidget->setCurrentWidget(ui->registerForm);

    //registerForm
    connect(ui->registerForm, &RegisterForm::TryToRegister,
        this, [this](QString username, QString password){
            serv->registerUser(username, password);
    });
    connect(serv, &ServerConnection::registerSuccessful, 
        this, [this](){
        serv->clearTokenFromUrl();
        ui->stackedWidget->setCurrentWidget(ui->login);
    });


    //login
    connect(login, &LoginForm::TryToLogin,
        this, [this](QString username, QString password){
            serv->login(username, password);
    });
    
    connect(serv, &ServerConnection::loginSuccessful, 
        this, [this](){
        ui->stackedWidget->setCurrentWidget(ui->touchpad);
        serv->wsHandle();
    });
    

    //touchpad
    connect(touchpad, &Touchpad::InputEvent,
        this, [this](QJsonObject event){
        serv->tryToReconnect();
        
        QJsonDocument doc(event);
        QString strJson(doc.toJson(QJsonDocument::Compact));
        qDebug() << "sending text message " << strJson;
        serv->sendTextMessage(strJson);
    });
    

}


MainWindow::~MainWindow(){
    delete ui;
}
