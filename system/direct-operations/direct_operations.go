package direct_operations

import (
	"app.modules/core"
	"app.modules/core/utils"
	"context"
	"encoding/json"
	"google.golang.org/api/option"
	"log"
	"os"
)


func ExitAllUsersAllRoom(clientOption option.ClientOption, ctx context.Context) {
	_system, err := core.NewSystem(ctx, clientOption)
	if err != nil {
		panic(err)
		return
	}
	
	_system.SendLiveChatMessage("全ユーザーを退室させます。", ctx)
	err = _system.ExitAllUserDefaultRoom(ctx)
	if err != nil {
		panic(err)
		return
	}
	err = _system.ExitAllUserNoSeatRoom(ctx)
	if err != nil {
		panic(err)
		return
	}
	_system.SendLiveChatMessage("全ユーザーを退室させました。", ctx)
}

func ExitSpecificUser(userId string, clientOption option.ClientOption, ctx context.Context) {
	_system, err := core.NewSystem(ctx, clientOption)
	if err != nil {
		panic(err)
		return
	}
	
	_system.SetProcessedUser(userId, "**")
	outCommandDetails := core.CommandDetails{
		CommandType:   core.Out,
		InOptions: core.InOptions{},
	}
	
	err = _system.Out(outCommandDetails, ctx)
	if err != nil {
		panic(err)
		return
	}
}

func ExportUsersCollectionJson(clientOption option.ClientOption, ctx context.Context) {
	_system, err := core.NewSystem(ctx, clientOption)
	if err != nil {
		panic(err)
		return
	}
	
	allUsersTotalStudySecList, err := _system.RetrieveAllUsersTotalStudySecList(ctx)
	if err != nil {
		panic(err)
		return
	}
	
	now := utils.JstNow()
	dateString := now.Format("2006-01-02_15-04-05")
	f, err := os.Create("./" + dateString + "_user-total-study-sec-list.json")
	if err != nil {
		panic(err)
		return
	}
	defer func() {_ = f.Close()}()
	
	jsonEnc := json.NewEncoder(f)
	jsonEnc.SetIndent("", "\t")
	err = jsonEnc.Encode(allUsersTotalStudySecList)
	if err != nil {
		panic(err)
		return
	}
	log.Println("finished exporting json.")
}
