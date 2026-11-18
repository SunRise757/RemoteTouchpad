#pragma once

#include <QWidget>

QT_BEGIN_NAMESPACE
namespace Ui {
    class RegisterForm;
}
QT_END_NAMESPACE


class RegisterForm: public QWidget
{
    Q_OBJECT

public:
    explicit RegisterForm(QWidget *parent = nullptr);
    
signals:
    void TryToRegister(QString username, QString password);

private:
    Ui::RegisterForm *ui;
};
