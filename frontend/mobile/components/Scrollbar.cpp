#include "Scrollbar.h"
#include "ui_Scrollbar.h"

#include <QMouseEvent>

Scrollbar::Scrollbar(QWidget *parent)
    : QWidget(parent)
    , ui(new Ui::Scrollbar)
{
    ui->setupUi(this);
    connect(ui->buttonScrollDown, &QPushButton::clicked,
        this, [this](){
            emit Scrolled(-1.0f);
    });
    connect(ui->buttonScrollUp, &QPushButton::clicked,
        this, [this](){
            emit Scrolled(1.0f);
    });
}

Scrollbar::~Scrollbar()
{
    delete ui;
}


void Scrollbar::mousePressEvent(QMouseEvent *event) {
    lastMousePosition = {0, 0};
    QWidget::mousePressEvent(event);
}

void Scrollbar::mouseReleaseEvent(QMouseEvent *event) {
    lastMousePosition = {0, 0};
    QWidget::mouseReleaseEvent(event);
}
void Scrollbar::mouseMoveEvent(QMouseEvent *event) {
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
    int sensitivity = 10;
    if (abs(delta.y())<sensitivity)
        return;

    emit Scrolled((event->pos() - lastMousePosition).y()/sensitivity);
    lastMousePosition = event->pos();
}