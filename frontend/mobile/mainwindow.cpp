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
            //emit ButtonClicked("buttonLeftClick");
            emit InputEvent(QJsonObject{
                {"type", "button"},
                {"name", "buttonLeftClick"}
            });
    });
    connect(ui->buttonRightClick, &QPushButton::clicked,
        this, [this](){
            //emit ButtonClicked("buttonRightClick");
            emit InputEvent(QJsonObject{
                {"type", "button"},
                {"name", "buttonRightClick"}
            });
    });
    connect(ui->buttonMiddleClick, &QPushButton::clicked,
        this, [this](){
            //emit ButtonClicked("buttonMiddleClick");
            emit InputEvent(QJsonObject{
                {"type", "button"},
                {"name", "buttonMiddleClick"}
            });
    });
    connect(ui->scrollbar, &Scrollbar::Scrolled,
        this, [this](float delta){
            emit InputEvent(QJsonObject{
                {"type", "scroll"},
                {"dy", delta}
            });
    });

    connect(ui->lineEdit, &QLineEdit::editingFinished,
        this, &MainWindow::textEdited);
}

MainWindow::~MainWindow(){
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
        emit InputEvent(QJsonObject{
            {"type", "button"},
            {"name", "buttonLeftClick"}
        });

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

    QPoint delta = event->pos() - lastMousePosition;
    if(delta.isNull()){
        return;
    }

    emit InputEvent(QJsonObject{
        {"type", "mouse"},
        {"dx", delta.x()},
        {"dy", delta.y()}
    });
    lastMousePosition = event->pos();
}

void MainWindow::textEdited(){
    QString text = ui->lineEdit->text();

    if(ui->lineEdit->text().isNull()){
        emit InputEvent(QJsonObject{
            {"type", "key"},
            {"key", "enter"},
            {"event", "press"}
        });
        return;
    }

    emit InputEvent(QJsonObject{
        {"type", "string"},
        {"string", text}
    });
    ui->lineEdit->setText("");

    return;

    if(text.startsWith(lastSavedString))
        text = text.last(text.length() - lastSavedString.length());

    /*
            if(lastSavedString.startsWith(text)){
                if(text.length() < lastSavedString.length()){
                    emit InputEvent(QJsonObject{
                        {"type", "key"},
                        {"key", "backspace"},
                        {"event", "press"}
                    });
                    lastSavedString = text;
                    return;
                }
            }
    */

        lastSavedString = text;

        emit InputEvent(QJsonObject{
            {"type", "string"},
            {"string", text}
        });
}
/*

void MainWindow::keyPressEvent(QKeyEvent *event) {

    if(event->key() == Qt::Key_Backspace) {
        emit InputEvent(QJsonObject{
            {"type", "key"},
            {"key", "backspace"},
            {"event", "press"}
        });
    }
}
void MainWindow::keyReleaseEvent(QKeyEvent *event) {

}
*/