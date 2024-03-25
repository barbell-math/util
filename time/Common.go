package time;

import "time"

// TODO - BAD, only has accuracy of a second, can delete?
// //After is the 'future' time, before is the 'past' time
// func Between(after time.Time, before time.Time) (func(t time.Time) bool) {
//     //if the 'past' time is after the 'future' time then switch them
//     if before.After(after) {
//         tmp:=before;
//         before=after;
//         after=tmp;
//     }
//     return func(t time.Time) bool {
//         return (before.AddDate(0, 0, -1).Before(t) &&
//             after.AddDate(0, 0, 1).After(t));
//     }
// }

// Gets the number of days between two dates, including negative days if before
// is not before before. ;)
func DaysBetween(before time.Time, after time.Time) int {
    if after.After(before) {
        return int(after.Sub(before).Hours()/24);
    } else {
        return -int(before.Sub(after).Hours()/24);
    }
}
