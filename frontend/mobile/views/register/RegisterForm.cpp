#include "RegisterForm.h"
#include "ui_RegisterForm.h"

RegisterForm::RegisterForm(QWidget *parent)
    : QWidget{parent}
    , ui(new Ui::RegisterForm)
{
    ui->setupUi(this);
    connect(ui->buttonRegister, &QPushButton::clicked,
        this, [this](){
            if (ui->lineEditPassword->text() == ui->lineEditRepeatPassword->text())
                emit TryToRegister(ui->lineEditUsername->text(), ui->lineEditPassword->text());
            else {
                ui->lineEditPassword->setText("");
                ui->lineEditRepeatPassword->setText("");
            }
        });
}

