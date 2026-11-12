#pragma once

#include <QDateTime>
#include <QMainWindow>
#include <QJsonObject>

#include "Touchpad.h"

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
    Touchpad *touchpad;
private:
    Ui::MainWindow *ui;
};
