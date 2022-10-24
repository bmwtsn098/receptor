package fuzz

import "strconv"
import "github.com/ansible/receptor/pkg/logger"
import "github.com/ansible/receptor/pkg/utils"
import "github.com/ansible/receptor/pkg/backends"
import "github.com/ansible/receptor/pkg/netceptor"

func mayhemit(bytes []byte) int {

    var num int
    if len(bytes) > 1 {
        num, _ = strconv.Atoi(string(bytes[0]))

        switch num {
    
        case 0:
            content := string(bytes)
            logger.GetLogLevelByName(content)
            return 0
    
        case 1:
            content := string(bytes)
            utils.TryFLock(content)
            return 0

        case 2:
            content := string(bytes)
            backends.NewUDPDialer(content, false)
            return 0

        case 3:
            var test backends.UDPDialerSession
            test.Send(bytes)
            return 0

        default:
            var test netceptor.PacketConn
            test.ReadFrom(bytes)
            return 0
        }
    }
    return 0
}

func Fuzz(data []byte) int {
    _ = mayhemit(data)
    return 0
}