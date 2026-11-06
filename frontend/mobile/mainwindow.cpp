#include "mainwindow.h"
#include "ui_mainwindow.h"
#include "components/Scrollbar.h"

#include <QMouseEvent>

MainWindow::MainWindow(QWidget *parent)
    : QMainWindow(parent)
    , ui(new Ui::MainWindow)
{
    ui->setupUi(this);
    connect(ui->buttonLeftClick, &QPushButton::clicked,
        this, [this](){
            emit ButtonClicked("buttonLeftClick");
    });
    connect(ui->buttonRightClick, &QPushButton::clicked,
        this, [this](){
            emit ButtonClicked("buttonRightClick");
    });
    connect(ui->buttonMiddleClick, &QPushButton::clicked,
        this, [this](){
            emit ButtonClicked("buttonMiddleClick");
    });
    connect(ui->scrollbar, &Scrollbar::Scrolled,
        this, &MainWindow::Scrolled);
}

MainWindow::~MainWindow()
{
    delete ui;
}

void MainWindow::mousePressEvent(QMouseEvent *event) {
    if(event->buttons().testFlag(Qt::LeftButton)) {
        lastMousePosition = {0, 0};
        lastMousePressTime = QDateTime::currentMSecsSinceEpoch();
        qDebug("Mouse press");
    }
    QMainWindow::mousePressEvent(event);
}

void MainWindow::mouseReleaseEvent(QMouseEvent *event) {
    lastMousePosition = {0, 0};

    qDebug("Mouse release");
    int clickTimeMs = 500;
    if (QDateTime::currentMSecsSinceEpoch() - lastMousePressTime < clickTimeMs)
        emit ButtonClicked("buttonLeftClick");

    QMainWindow::mouseReleaseEvent(event);
}


void MainWindow::mouseMoveEvent(QMouseEvent *event) {

    event->accept();

    if(!event->buttons().testFlag(Qt::LeftButton)){
        lastMousePosition = {0, 0};
        return;
    }

    if (lastMousePosition.isNull()){
        lastMousePosition = event->pos();
        return;
    }

    QPoint delta = lastMousePosition - event->pos();
    if(delta.isNull()){
        return;
    }

    emit TouchpadMoved(event->pos() - lastMousePosition);
    lastMousePosition = event->pos();
}
