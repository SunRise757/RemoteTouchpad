#include "mainwindow.h"
#include "ui_mainwindow.h"
#include "components/Scrollbar.h"

#include <QMouseEvent>

MainWindow::MainWindow(QWidget *parent)
    : QMainWindow(parent)
    , ui(new Ui::MainWindow)
{
    ui->setupUi(this);

}

MainWindow::~MainWindow()
{
    delete ui;
}
void MainWindow::mouseMoveEvent(QMouseEvent *event) {
    QMainWindow::mouseMoveEvent(event);
    
    qDebug() << "Mouse moved:" << event->pos();
    emit TouchpadMoved();
}
