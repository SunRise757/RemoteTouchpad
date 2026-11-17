#include "Touchpad.h"
#include "ui_Touchpad.h"
#include "components/Scrollbar.h"

#include <QMouseEvent>

Touchpad::Touchpad(QWidget *parent)
    : QWidget(parent)
    , ui(new Ui::Touchpad)
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
        this, &Touchpad::textEdited);
}

Touchpad::~Touchpad(){
    delete ui;
}

void Touchpad::mousePressEvent(QMouseEvent *event) {
    if(event->buttons().testFlag(Qt::LeftButton)) {
        lastMousePosition = {0, 0};
        lastMousePressTime = QDateTime::currentMSecsSinceEpoch();
    }
    QWidget::mousePressEvent(event);
}

void Touchpad::mouseReleaseEvent(QMouseEvent *event) {
    lastMousePosition = {0, 0};

    int clickTimeMs = 500;
    if (QDateTime::currentMSecsSinceEpoch() - lastMousePressTime < clickTimeMs)
        emit InputEvent(QJsonObject{
            {"type", "button"},
            {"name", "buttonLeftClick"}
        });

    QWidget::mouseReleaseEvent(event);
}


void Touchpad::mouseMoveEvent(QMouseEvent *event) {

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

void Touchpad::textEdited(){
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

        lastSavedString = text;

        emit InputEvent(QJsonObject{
            {"type", "string"},
            {"string", text}
        });
}
