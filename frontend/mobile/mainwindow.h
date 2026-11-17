#pragma once

#include <QDateTime>
#include <QMainWindow>
#include <QJsonObject>

#include "Touchpad.h"
#include "LoginForm.h"
#include "serverconnection.h"

QT_BEGIN_NAMESPACE
namespace Ui {
    class MainWindow;
}
QT_END_NAMESPACE

class MainWindow : public QMainWindow
{
    Q_OBJECT

public:
    explicit MainWindow(QWidget *parent = nullptr);
    ~MainWindow();
    Touchpad *touchpad;
    LoginForm *login;
private:
    Ui::MainWindow *ui;
    ServerConnection *serv;
};
