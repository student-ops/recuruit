package session

//session sturcuture
//session[userid] = 1 or 0
//session[userid] = tmp_thread_id
var Tmp_session map[int]int
var Tmp_thread_session map[int]int