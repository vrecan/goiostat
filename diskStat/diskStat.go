package diskStat
import (
   "strconv"
   )
type DiskStat struct {
	id int64
	partId int64
	device string
	readsCompleted int64
	readsMerged int64
	sectorsRead int64
	millisReading int64
	writesCompleted int64
	writesMerged int64
	sectorsWritten int64
	millisWriting int64
	ioInProgress int64
	millisDoingIo int64
	weightedMillisDoingIo int64
}

/*
  Simple function that takes a string and converts it to a stat
  TODO: Figure out better way to map string to struct
*/
func LineToStat(line []string) (stat DiskStat, err error){
	var tmp int64
	tmp,err = strconv.ParseInt(line[0], 10, 64)
	if nil != err {
		return stat, err;
	}
	stat.id = tmp

	tmp,err = strconv.ParseInt(line[1], 10, 64)
	if nil != err {
		return stat, err;
	}
	stat.partId = tmp

  	stat.device = line[2]


	tmp,err = strconv.ParseInt(line[3], 10, 64)
	if nil != err {
		return stat, err;
	}
	stat.readsCompleted = tmp

	tmp,err = strconv.ParseInt(line[4], 10, 64)
	if nil != err {
		return stat, err;
	}
	stat.readsMerged = tmp		


	tmp,err = strconv.ParseInt(line[5], 10, 64)
	if nil != err {
		return stat, err;
	}
	stat.sectorsRead = tmp

	tmp,err = strconv.ParseInt(line[6], 10, 64)
	if nil != err {
		return stat, err;
	}
	stat.millisReading = tmp	

	tmp,err = strconv.ParseInt(line[7], 10, 64)
	if nil != err {
		return stat, err;
	}
	stat.writesCompleted = tmp		
	
	tmp,err = strconv.ParseInt(line[8], 10, 64)
	if nil != err {
		return stat, err;
	}
	stat.writesMerged = tmp	

	tmp,err = strconv.ParseInt(line[9], 10, 64)
	if nil != err {
		return stat, err;
	}
	stat.sectorsWritten = tmp		

	tmp,err = strconv.ParseInt(line[10], 10, 64)
	if nil != err {
		return stat, err;
	}
	stat.millisWriting = tmp			

	tmp,err = strconv.ParseInt(line[11], 10, 64)
	if nil != err {
		return stat, err;
	}
	stat.ioInProgress = tmp	

	tmp,err = strconv.ParseInt(line[12], 10, 64)
	if nil != err {
		return stat, err;
	}
	stat.millisDoingIo = tmp	

    tmp,err = strconv.ParseInt(line[13], 10, 64)
	if nil != err {
		return stat, err;
	}
	stat.weightedMillisDoingIo = tmp	


  	return stat, err;

}