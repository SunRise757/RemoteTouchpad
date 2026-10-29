#include "Scrollbar.h"
#include "ui_Scrollbar.h"

#include <QMouseEvent>

Scrollbar::Scrollbar(QWidget *parent)
    : QWidget(parent)
    , ui(new Ui::Scrollbar)
{
    ui->setupUi(this);
}

Scrollbar::~Scrollbar()
{
    delete ui;
}
void Scrollbar::mouseMoveEvent(QMouseEvent *event) {
    QWidget::mouseMoveEvent(event);
    qDebug() << "Scrollbar moved:" << event->pos();
    emit Scrolled();
    event->accept();
}