#include "mainwindow.h"
#include "ui_mainwindow.h"
#include <QMouseEvent>

MainWindow::MainWindow(QWidget *parent)
    : QMainWindow(parent)
    , ui(new Ui::MainWindow)
{
    ui->setupUi(this);
    /*
    connect(ui->buttonTouchpad, &QPushButton::mouseDoubleClickEvent,
    this, [](){});
    */
    //ui->buttonLeftClick->setObjectName->name()
}

MainWindow::~MainWindow()
{
    delete ui;
}
void MainWindow::mouseMoveEvent(QMouseEvent *event) {
    QMainWindow::mouseMoveEvent(event);
    qDebug() << "Mouse moved:" << event->pos();
    emit onTouchpadMoved();
}
