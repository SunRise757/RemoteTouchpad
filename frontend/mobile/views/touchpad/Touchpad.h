#pragma once

#include <QWidget>
#include <QDateTime>
#include <QJsonObject>

#include "Scrollbar.h"

QT_BEGIN_NAMESPACE
namespace Ui {
    class Touchpad;
}
QT_END_NAMESPACE

class Touchpad : public QWidget
{
    Q_OBJECT

public:
    explicit Touchpad(QWidget *parent = nullptr);
    ~Touchpad();
protected:
    void mouseMoveEvent(QMouseEvent *event) override;
    void mousePressEvent(QMouseEvent *event) override;
    void mouseReleaseEvent(QMouseEvent *event) override;
    /*
    void keyPressEvent(QKeyEvent *event) override;
    void keyReleaseEvent(QKeyEvent *event) override;
    */
    void textEdited();
signals:
    /*void TouchpadMoved(QPoint delta);
    void Scrolled(float delta);
    void ButtonClicked(QString name);
    */
    void InputEvent(QJsonObject event);
private:
    QString lastSavedString;
    QPoint lastMousePosition{0, 0};
    qint64 lastMousePressTime;
    Ui::Touchpad *ui;
};
