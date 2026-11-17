#include "LoginForm.h"
#include "ui_LoginForm.h"

LoginForm::LoginForm(QWidget *parent)
    : QWidget{parent}
    , ui(new Ui::LoginForm)
{
    ui->setupUi(this);
    connect(ui->buttonLogin, &QPushButton::clicked,
        this, [this](){
            emit TryToLogin(ui->lineEditUsername->text(), ui->lineEditPassword->text());
        });
}

