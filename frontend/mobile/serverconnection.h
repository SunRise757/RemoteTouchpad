#pragma once

#include <QWebSocket>
#include <QNetworkAccessManager>
#include <QJsonObject>
#include <QJsonDocument>

#include <QObject>

class ServerConnection : public QObject
{
    Q_OBJECT
public:
    explicit ServerConnection(QObject *parent = nullptr);
    void tryToReconnect();
    void sendTextMessage(QString message) const;
    void registerUser(QString username, QString password);
    void login(QString username, QString password);
    void logout(QString username);
    void protectedRoute();
    void wsHandle();
    void clearTokenFromUrl() const;
    QString registerToken;

signals:
    void loginSuccessful();
    void registerSuccessful();

private:
    QNetworkAccessManager *manager;
    QWebSocket *socket;
    QUrl lastConnectedUrl = getWebsocketUrl();
    QUrl getWebsocketUrl();
    QUrl getUrl() const;
    QString getRegisterToken() const;
    qint64 lastReconnectTime;
};
