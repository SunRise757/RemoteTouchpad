#include <QApplication>
#include <QJsonObject>
#include <QJsonDocument>

#include "serverconnection.h"
#include "mainwindow.h"

int main(int argc, char *argv[])
{
    QApplication a(argc, argv);

    MainWindow w;
    w.show();

    ServerConnection serv;

    QObject::connect(w.touchpad, &Touchpad::InputEvent,
        &w, [serv](QJsonObject event){
        serv.tryToReconnect();

        QJsonDocument doc(event);
        QString strJson(doc.toJson(QJsonDocument::Compact));
        serv.sendTextMessage(strJson);
    });

    return a.exec();
}
