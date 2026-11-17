#include <QDateTime>
#include <QNetworkRequest>
#include <QNetworkReply>
#include <QUrlQuery>
#include <emscripten.h>
#include <emscripten/val.h>
#include <QByteArray>
#include <cstdlib>

#include "serverconnection.h"

//defines getCsrfToken function
EM_JS(char*, getCsrfToken, (), {
    const prefix = "csrf_token=";
    const entry = document.cookie
        .split("; ")
        .find(cookie => cookie.startsWith(prefix));

    if (!entry) return 0;
    return stringToNewUTF8(
        decodeURIComponent(entry.substring(prefix.length))
    );
});
QByteArray csrfToken()
{
    char *value = getCsrfToken();
    if (!value) return {};
    QByteArray result(value);
    std::free(value);
    return result;
}


ServerConnection::ServerConnection(QObject *parent)
    : QObject(parent)
    , manager(new QNetworkAccessManager(this))
    , socket(new QWebSocket())
    , lastReconnectTime(0)
    , lastConnectedUrl(getWebsocketUrl())
{
    
}

QUrl ServerConnection::getWebsocketUrl(){
    const auto location = emscripten::val::global("location");
    const auto protocol = location["protocol"].as<std::string>();
    const auto host = location["host"].as<std::string>();

    const std::string wsScheme = protocol == "https:" ? "wss://" : "ws://";
    return QUrl(QString::fromStdString(wsScheme + host + "/ws"));
}
QUrl ServerConnection::getUrl() const{
    const auto location = emscripten::val::global("location");
    const auto protocol = location["protocol"].as<std::string>();
    const auto host = location["host"].as<std::string>();
    
    return QUrl(QString::fromStdString(protocol + "//" + host));
}

void ServerConnection::registerUser(QString username, QString password){

}

void ServerConnection::login(QString username, QString password) {
    QUrl url = getUrl();
    url.setPath("/login");
    qDebug() << "url = " << url;
    QNetworkRequest request(url);

    request.setHeader(
        QNetworkRequest::ContentTypeHeader,
        "application/x-www-form-urlencoded"
    );
    
    QUrlQuery form;
    form.addQueryItem("username", username);
    form.addQueryItem("password", password);
    
    QJsonObject data = QJsonObject{
        {"username", username}, 
        {"password", password}
    };
    
    QNetworkReply *reply = manager->post(
        request, 
        form.query(QUrl::FullyEncoded).toUtf8()
    );
    
    connect(reply, &QNetworkReply::finished, 
        this, [this, reply]() {
        
        if (reply->error() == QNetworkReply::NoError) {
            
            QByteArray replyData = reply->readAll();
            qDebug() << replyData;

            QJsonParseError js;
            QJsonObject data = QJsonDocument::fromJson(replyData, &js).object();
            
            if(js.error){
                qDebug() << "Error parsing JSON:" << js.errorString();
            }
            
            if (data["success"].toBool()){
                qDebug() << "Login successful";
                emit loginSuccessful();
            }
            
        } else {
            qWarning() << reply->errorString();
        }
        reply->deleteLater();
    });
    

}

void ServerConnection::protectedRoute(){
    
    //this is a test
    QUrl url = getUrl();
    url.setPath("/protected");
    //url.setPath("/ws");
    qDebug() << "url = " << url;
    QNetworkRequest request(url);

    request.setHeader(
        QNetworkRequest::ContentTypeHeader,
        "application/x-www-form-urlencoded"
    );
    request.setRawHeader("X-CSRF-Token", csrfToken());
    
    QUrlQuery form;
    form.addQueryItem("username", "admin");
    //form.addQueryItem("password", password);
    
    QJsonObject data = QJsonObject{
        {"username", "admin"}//, 
        //{"password", password}
    };
    
    QNetworkReply *reply = manager->post(
        request, 
        form.query(QUrl::FullyEncoded).toUtf8()
    );
    
    connect(reply, &QNetworkReply::finished, 
        this, [this, reply]() {
        
        
        if (reply->error() == QNetworkReply::NoError) {
            
            QByteArray replyData = reply->readAll();
            qDebug() << replyData;

            QJsonParseError js;
            QJsonObject data = QJsonDocument::fromJson(replyData, &js).object();
            
            if(js.error){
                qDebug() << "Error parsing JSON:" << js.errorString();
            }
            
        } else {
            qWarning() << reply->errorString();
        }
        reply->deleteLater();
    });
}

void ServerConnection::logout(QString username){

}

void ServerConnection::wsHandle(){
    
    QObject::connect(socket, &QWebSocket::connected,
        [this](){
        qDebug("Socket connected");
        socket->sendTextMessage("Hello from client!");
    });

    QObject::connect(socket, &QWebSocket::textMessageReceived,
        [this](QString message) {
        qDebug() << "Message received:" << message;
    });

    QObject::connect(socket, &QWebSocket::errorOccurred,
        [this](QAbstractSocket::SocketError error) {
        qDebug() << "Socket error:" << this->socket->errorString();
    });

    auto disconnection = QObject::connect(socket, &QWebSocket::disconnected,
        [this](){
            tryToReconnect();
    });

    lastConnectedUrl = getWebsocketUrl();
    socket->open(lastConnectedUrl);
}

void ServerConnection::tryToReconnect() const{
void ServerConnection::tryToReconnect(){
    if (socket->state() == QAbstractSocket::ConnectedState && socket->isValid())
        return;
    
    if(QDateTime::currentSecsSinceEpoch() < lastReconnectTime + 2){
        qDebug() << "reconnect spam delay";
        return;
    }
    qDebug() << "trying to reconnect";
    socket->open(lastConnectedUrl);
}

void ServerConnection::sendTextMessage(QString message) const{
    if(!socket->isValid()){
        qDebug() << "socket invalid";
        return;
    }
    if(socket->state() != QAbstractSocket::ConnectedState) {
        qDebug() << "socket connected state = " << socket->state();    
        return;
    }
        socket->sendTextMessage(message);
    
}