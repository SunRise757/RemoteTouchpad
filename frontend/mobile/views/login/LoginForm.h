#pragma once

#include <QWidget>

QT_BEGIN_NAMESPACE
namespace Ui {
    class LoginForm;
}
QT_END_NAMESPACE


class LoginForm: public QWidget
{
    Q_OBJECT

public:
    explicit LoginForm(QWidget *parent = nullptr);
    
signals:
    void TryToLogin(QString username, QString password);

private:
    Ui::LoginForm *ui;
};
