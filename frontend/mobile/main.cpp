#include <QApplication>
#include <QWebSocket>
#include <emscripten/val.h>
#include <string>
#include "mainwindow.h"

QUrl getWebsocketUrl(){
    const auto location = emscripten::val::global("location");
    const auto protocol = location["protocol"].as<std::string>();
    const auto host = location["host"].as<std::string>();

    const std::string wsScheme = protocol == "https:" ? "wss://" : "ws://";
    return QUrl(QString::fromStdString(wsScheme + host + "/ws"));
}

int main(int argc, char *argv[])
{
    QApplication a(argc, argv);

    MainWindow w;
    w.show();

    QWebSocket *socket = new QWebSocket;
//web socket connections
    QObject::connect(socket, &QWebSocket::connected, [socket](){
        qDebug("Socket connected");
        socket->sendTextMessage("Hello from client!");
    });

    QObject::connect(socket, &QWebSocket::textMessageReceived, [socket](QString message) {
        qDebug() << "Message received:" << message;
    });

    QObject::connect(socket, &QWebSocket::errorOccurred, [socket](QAbstractSocket::SocketError error) {
        qDebug() << "Socket error:" << socket->errorString();
    });

    return a.exec();
}
