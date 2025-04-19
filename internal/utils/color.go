package utils

import "fmt"

/*

switch level {
	case LEVEL_TRACE:
		colorTimeName = fmt.Sprintf("\033[38m%v %s %s", strTimeFmt, strPID, Name)
	case LEVEL_DEBUG:
		colorTimeName = fmt.Sprintf("\033[34m%v %s %s", strTimeFmt, strPID, Name)
	case LEVEL_INFO:
		colorTimeName = fmt.Sprintf("\033[32m%v %s %s", strTimeFmt, strPID, Name)
	case LEVEL_WARN:
		colorTimeName = fmt.Sprintf("\033[33m%v %s %s", strTimeFmt, strPID, Name)
	case LEVEL_ERROR:
		colorTimeName = fmt.Sprintf("\033[31m%v %s %s", strTimeFmt, strPID, Name)
	case LEVEL_FATAL:
		colorTimeName = fmt.Sprintf("\033[35m%v %s %s", strTimeFmt, strPID, Name)
	case LEVEL_PANIC:
		colorTimeName = fmt.Sprintf("\033[35m%v %s %s", strTimeFmt, strPID, Name)
	default:
		colorTimeName = fmt.Sprintf("\033[34m%v %s %s", strTimeFmt, strPID, Name)
	}
*/

func White(s any) string {
	return fmt.Sprintf("\033[38m%v\033[0m", s)
}

func Green(s any) string {
	return fmt.Sprintf("\033[32m%v\033[0m", s)
}

func Red(s any) string {
	return fmt.Sprintf("\033[31m%v\033[0m", s)
}
