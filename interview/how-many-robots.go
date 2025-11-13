package main

// interview

/*
Есть последовательность запросов пользователей, каждый запрос –
это пара (time, userId), запросы всегда приходят в отсортированном
по времени порядке.
Нужно уметь быстро отвечать: "сколько за последние k секунд было
пользователей, которые задали >= limit запросов". k > 0, limit > 0.
*/

// All methods are guaranteed to be called with non-decreasing (>=) time

//type UserStatistics struct {
//	Window int
//	Limit  int
//}

//func NewUserStatistics(window time.Duration, limit int) *UserStatistics {
//	return &UserStatistics{Window: window, Limit: limit}
//}

// AddEvent registers one click of user.
//func (u *UserStatistics) AddEvent(now time.Time, userID int) {
//}

// GetRobotCount returns number of robots.
// User is robot at given time now <=> they have >= limit clicks since (>=) now - window.
//func (u *UserStatistics) GetRobotCount(now time.Time) int {
//	return 0
//}

//

type intTime = int

type Event struct {
	userID    int
	eventTime intTime
}

type UserStatistics struct {
	Window   intTime
	Limit    int
	Events   []Event
	ReqCnt   map[int]int
	RobotCnt int
}

func NewUserStatistics(window intTime, limit int) *UserStatistics {
	return &UserStatistics{Window: window, Limit: limit, Events: nil, ReqCnt: make(map[int]int), RobotCnt: 0}
}

func (u *UserStatistics) AddEvent(now intTime, userID int) {
	u.Events = append(u.Events, Event{userID, now})
	u.ReqCnt[userID]++
	if u.ReqCnt[userID] == u.Limit {
		u.RobotCnt++
	}
}

func (u *UserStatistics) GetRobotCount(now intTime) int {
	// сначала бежим по всем значениям, секунда которых меньше нас - окно, кикаем из списка и из мапы
	for i := 0; i < len(u.Events); i++ {
		event := u.Events[i]
		if event.eventTime >= (now - u.Window) {
			break
		}
		u.ReqCnt[event.userID]--
		u.Events = u.Events[1:]
		if u.ReqCnt[event.userID]+1 == u.Limit {
			u.RobotCnt--
		}
	}
	return u.RobotCnt
}

func main() {
	stats := NewUserStatistics(10, 3)
	stats.AddEvent(0, 1)
	stats.AddEvent(3, 2)
	stats.AddEvent(4, 3)
	stats.AddEvent(4, 2)
	stats.AddEvent(6, 1)
	stats.AddEvent(7, 1)
	stats.GetRobotCount(8)  // есть 1
	stats.GetRobotCount(11) // нету никаво
	stats.AddEvent(11, 2)
	stats.GetRobotCount(12) // снова есть
}
