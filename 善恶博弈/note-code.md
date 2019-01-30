# 代码结构

    PlayerFactory
        Player
            BasicPlayer.go  
            Player1.go
            Player2.go
            ...
        Strategy
            Strategy.go
        PlayerFactory.go
    Entrance.go

## PlayerFactory 玩家创建简单工厂

通过策略接口将各种不同的策略实例化为一个玩家的简单工厂

## Player 策略包

保存各种策略的类

## BasicPlayer.go 和 xxx.go

策略基类，xxx.go 如Player1.go表示策略1

## Strategy 和 Strategy.go

策略接口，用以实现策略模式

## Entrance.go

程序入口