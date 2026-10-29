#pragma once

#include <QWidget>

QT_BEGIN_NAMESPACE
namespace Ui {
    class Scrollbar;
}
QT_END_NAMESPACE

class Scrollbar: public QWidget
{
    Q_OBJECT

public:
    explicit Scrollbar(QWidget *parent = nullptr);
    ~Scrollbar();
protected:
    void mouseMoveEvent(QMouseEvent *event) override;

signals:
    float Scrolled();

private:
    Ui::Scrollbar *ui;
};
