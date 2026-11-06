#pragma once

#include <QDateTime>
#include <QMainWindow>
#include "Scrollbar.h"

QT_BEGIN_NAMESPACE
namespace Ui {
    class MainWindow;
}
QT_END_NAMESPACE

class MainWindow : public QMainWindow
{
    Q_OBJECT

public:
    explicit MainWindow(QWidget *parent = nullptr);
    ~MainWindow();
protected:
    void mouseMoveEvent(QMouseEvent *event) override;
    void mousePressEvent(QMouseEvent *event) override;
    void mouseReleaseEvent(QMouseEvent *event) override;

signals:
    void TouchpadMoved(QPoint delta);
    void Scrolled(float delta);
    void ButtonClicked(QString name);
private:
    QPoint lastMousePosition{0, 0};
    qint64 lastMousePressTime;
    Ui::MainWindow *ui;
};
