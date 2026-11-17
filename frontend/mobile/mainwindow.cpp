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
    ui->stackedWidget->setCurrentWidget(ui->login);

    serv = new ServerConnection(this);
    //login
    connect(login, &LoginForm::TryToLogin,
        this, [this](QString username, QString password){
            serv->login(username, password);
    });
    
    connect(serv, &ServerConnection::loginSuccessful, 
        this, [this](){
        ui->stackedWidget->setCurrentWidget(ui->touchpad);
        serv->wsHandle();
        //test bullshit deletre later
        //serv->protectedRoute();
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
