package main

import (
	//	"fmt"
	log "github.com/Sirupsen/logrus"

	//"bitbucket.org/crmsib/parse_gov/contracts"
	//"bitbucket.org/crmsib/parse_gov/notifications"
	"bitbucket.org/crmsib/parser_gov/nsi"
	"bitbucket.org/crmsib/parser_gov/notifications"
	//"runtime"
	"database/sql"
	"sync/atomic"

	"errors"
	"fmt"
	"strings"
)

func ignore(t string) {
	log.Info("ignore:", t)
	// Здесь должна быть обертка для логов.
}

// ParserWorker - control switch of xml parsing
func ParserWorker(db *sql.DB, xmlName *XmlFile, flagIsSave int, counters *Stats) error {
	var count int32
	//var countEF int32
	x := xmlName

	var err error
	log.Info(x.Name, "In PROCESS!\n")
	switch x.Type {
		case "fcs":
			ignore("fcs")
			log.Println("ignore:", string(x.Bucket()), x.Name)
			err = nil
			break

		// Справочники
		case "nsiOrganizationType":
			err = nsi.ParseNsiOrgType(x.Data, x.Name, db)
			if err != nil {
				atomic.AddInt32(&counters.withErrors, 1)
				atomic.AddInt32(&counters.nsiOrganizationTypewithErrors, 1)
			} else {
				atomic.AddInt32(&counters.nsi, 1)
				atomic.AddInt32(&counters.nsiOrganizationType, 1)
			}
			break
		case "nsiOrganizationTypesList":
			err = nsi.ParseNsiOrgType(x.Data, x.Name, db)
			if err != nil {
				atomic.AddInt32(&counters.withErrors, 1)
				atomic.AddInt32(&counters.nsiOrganizationTypewithErrors, 1)
			} else {
				atomic.AddInt32(&counters.nsi, 1)
				atomic.AddInt32(&counters.nsiOrganizationType, 1)
			}
			break
		case "nsiOrganization":
			err = nsi.ParseNsiOrg(x.Data, x.Name, db)
			if err != nil {
				atomic.AddInt32(&counters.withErrors, 1)
				atomic.AddInt32(&counters.nsiOrganizationwithErrors, 1)
			} else {
				atomic.AddInt32(&counters.nsi, 1)
				atomic.AddInt32(&counters.nsiOrganization, 1)
			}
			break
		case "nsiOrganizationList":
			err = nsi.ParseNsiOrg(x.Data, x.Name, db)
			if err != nil {
				atomic.AddInt32(&counters.withErrors, 1)
				atomic.AddInt32(&counters.nsiOrganizationwithErrors, 1)
			} else {
				atomic.AddInt32(&counters.nsi, 1)
				atomic.AddInt32(&counters.nsiOrganization, 1)
			}
			break
		case "nsiBudget":
			err = nsi.ParseNsiBudget(x.Data, x.Name, db)
			if err != nil {
				atomic.AddInt32(&counters.withErrors, 1)
				atomic.AddInt32(&counters.nsiBudgetwithErrors, 1)
			} else {
				atomic.AddInt32(&counters.nsi, 1)
				atomic.AddInt32(&counters.nsiBudget, 1)
			}
			break
		case "nsiBudgetList":
			err = nsi.ParseNsiBudget(x.Data, x.Name, db)
			if err != nil {
				atomic.AddInt32(&counters.withErrors, 1)
				atomic.AddInt32(&counters.nsiBudgetwithErrors, 1)
			} else {
				atomic.AddInt32(&counters.nsi, 1)
				atomic.AddInt32(&counters.nsiBudget, 1)
			}
			break
		case "nsiOKEI":
			err = nsi.ParseNsiOKEI(x.Data, x.Name, db)
			if err != nil {
				atomic.AddInt32(&counters.withErrors, 1)
				atomic.AddInt32(&counters.nsiOKEIwithErrors, 1)
			} else {
				atomic.AddInt32(&counters.nsi, 1)
				atomic.AddInt32(&counters.nsiOKEI, 1)
			}
			break
		case "nsiOKEIList":
			err = nsi.ParseNsiOKEI(x.Data, x.Name, db)
			if err != nil {
				atomic.AddInt32(&counters.withErrors, 1)
				atomic.AddInt32(&counters.nsiOKEIwithErrors, 1)
			} else {
				atomic.AddInt32(&counters.nsi, 1)
				atomic.AddInt32(&counters.nsiOKEI, 1)
			}
			break
		case "nsiOKPD":
			err = nsi.ParseNsiOKPD(x.Data, x.Name, db)
			if err != nil {
				atomic.AddInt32(&counters.withErrors, 1)
				atomic.AddInt32(&counters.nsiOKPDwithErrors, 1)
			} else {
				atomic.AddInt32(&counters.nsi, 1)
				atomic.AddInt32(&counters.nsiOKPD, 1)
			}
			break
		case "nsiOKPDList":
			err = nsi.ParseNsiOKPD(x.Data, x.Name, db)
			if err != nil {
				atomic.AddInt32(&counters.withErrors, 1)
				atomic.AddInt32(&counters.nsiOKPDwithErrors, 1)
			} else {
				atomic.AddInt32(&counters.nsi, 1)
				atomic.AddInt32(&counters.nsiOKPD, 1)
			}
			break
		case "nsiOKPD2":
			err = nsi.ParseNsiOKPD(x.Data, x.Name, db)
			if err != nil {
				atomic.AddInt32(&counters.withErrors, 1)
				atomic.AddInt32(&counters.nsiOKPD2withErrors, 1)
			} else {
				atomic.AddInt32(&counters.nsi, 1)
				atomic.AddInt32(&counters.nsiOKPD2, 1)
			}
			break
		case "nsiOKPD2List":
			err = nsi.ParseNsiOKPD(x.Data, x.Name, db)
			if err != nil {
				atomic.AddInt32(&counters.withErrors, 1)
				atomic.AddInt32(&counters.nsiOKPD2withErrors, 1)
			} else {
				atomic.AddInt32(&counters.nsi, 1)
				atomic.AddInt32(&counters.nsiOKPD2, 1)
			}
			break
		case "nsiOKTMO":
			err = nsi.ParseNsiOKTMO(x.Data, x.Name, db)
			if err != nil {
				atomic.AddInt32(&counters.withErrors, 1)
				atomic.AddInt32(&counters.nsiOKTMOwithErrors, 1)
			} else {
				atomic.AddInt32(&counters.nsi, 1)
				atomic.AddInt32(&counters.nsiOKTMO, 1)
			}
			break
		case "nsiOKTMOList":
			err = nsi.ParseNsiOKTMO(x.Data, x.Name, db)
			if err != nil {
				atomic.AddInt32(&counters.withErrors, 1)
				atomic.AddInt32(&counters.nsiOKTMOwithErrors, 1)
			} else {
				atomic.AddInt32(&counters.nsi, 1)
				atomic.AddInt32(&counters.nsiOKTMO, 1)
			}
			break
		case "nsiOffBudget":
			err = nsi.ParseNsiOffBudget(x.Data, x.Name, db)
			if err != nil {
				atomic.AddInt32(&counters.withErrors, 1)
				atomic.AddInt32(&counters.nsiOffBudgetwithErrors, 1)
			} else {
				atomic.AddInt32(&counters.nsi, 1)
				atomic.AddInt32(&counters.nsiOffBudget, 1)
			}
			break
		case "nsiOffBudgetList":
			err = nsi.ParseNsiOffBudget(x.Data, x.Name, db)
			if err != nil {
				atomic.AddInt32(&counters.withErrors, 1)
				atomic.AddInt32(&counters.nsiOffBudgetwithErrors, 1)
			} else {
				atomic.AddInt32(&counters.nsi, 1)
				atomic.AddInt32(&counters.nsiOffBudget, 1)
			}
			break
		case "nsiOffBudgetType":
			err = nsi.ParseNsiOffBudget(x.Data, x.Name, db)
			if err != nil {
				atomic.AddInt32(&counters.withErrors, 1)
				atomic.AddInt32(&counters.nsiOffBudgetTypewithErrors, 1)
			} else {
				atomic.AddInt32(&counters.nsi, 1)
				atomic.AddInt32(&counters.nsiOffBudgetType, 1)
			}
			break
		case "nsiPurchasePreference":
			err = nsi.ParseNsiPurchasePreference(x.Data, x.Name, db)
			if err != nil {
				atomic.AddInt32(&counters.withErrors, 1)
				atomic.AddInt32(&counters.nsiPurchasePreferencewithErrors, 1)
			} else {
				atomic.AddInt32(&counters.nsi, 1)
				atomic.AddInt32(&counters.nsiPurchasePreference, 1)
			}
			break
		case "nsiPurchasePreferenceList":
			err = nsi.ParseNsiPurchasePreference(x.Data, x.Name, db)
			if err != nil {
				atomic.AddInt32(&counters.withErrors, 1)
				atomic.AddInt32(&counters.nsiPurchasePreferencewithErrors, 1)
			} else {
				atomic.AddInt32(&counters.nsi, 1)
				atomic.AddInt32(&counters.nsiPurchasePreference, 1)
			}
			break
		case "nsiETPs":
			err = nsi.ParseNsiETP(x.Data, x.Name, db)
			if err != nil {
				atomic.AddInt32(&counters.withErrors, 1)
				atomic.AddInt32(&counters.nsiETPswithErrors, 1)
			} else {
				atomic.AddInt32(&counters.nsi, 1)
				atomic.AddInt32(&counters.nsiETPs, 1)
			}
			break
		case "nsiPlacingWayList":
			err = nsi.ParseNsiPlacingWay(x.Data, x.Name, db)
			if err != nil {
				atomic.AddInt32(&counters.withErrors, 1)
				atomic.AddInt32(&counters.nsiPlacingWayListwithErrors, 1)
			} else {
				atomic.AddInt32(&counters.nsi, 1)
				atomic.AddInt32(&counters.nsiPlacingWayList, 1)
			}
			break
		case "nsiFarmDrugsDictionary":
			err = nsi.ParseNsiFarmDrugsDic(x.Data, x.Name, db)
			if err != nil {
				atomic.AddInt32(&counters.withErrors, 1)
				atomic.AddInt32(&counters.nsiFarmDrugsDictionarywithErrors, 1)
			} else {
				atomic.AddInt32(&counters.nsi, 1)
				atomic.AddInt32(&counters.nsiFarmDrugsDictionary, 1)
			}
			break

		// Аукционы
		case "fcsNotificationEA44":
			err = notifications.ParseNotification(x.Data, db)
			if err != nil {
				atomic.AddInt32(&counters.withErrors, 1)
				atomic.AddInt32(&counters.notificationEFwithErrors, 1)
			} else {
				atomic.AddInt32(&counters.notifications, 1)
				atomic.AddInt32(&counters.notificationEF, 1)
			}
			break
		case "fcsNotificationEP44":
			err = notifications.ParseNotification(x.Data, db)
			if err != nil {
				atomic.AddInt32(&counters.withErrors, 1)
				atomic.AddInt32(&counters.notificationEPwithErrors, 1)
			} else {
				atomic.AddInt32(&counters.notifications, 1)
				atomic.AddInt32(&counters.notificationEP, 1)
			}
			break
		case "fcsNotificationOKD44":
			err = notifications.ParseNotification(x.Data, db)
			if err != nil {
				atomic.AddInt32(&counters.withErrors, 1)
				atomic.AddInt32(&counters.notificationOKDwithErrors, 1)
			} else {
				atomic.AddInt32(&counters.notifications, 1)
				atomic.AddInt32(&counters.notificationOKD, 1)
			}
			break
		case "fcsNotificationOKU44":
			err = notifications.ParseNotification(x.Data, db)
			if err != nil {
				atomic.AddInt32(&counters.withErrors, 1)
				atomic.AddInt32(&counters.notificationOKUwithErrors, 1)
			} else {
				atomic.AddInt32(&counters.notifications, 1)
				atomic.AddInt32(&counters.notificationOKU, 1)
			}
			break
		case "fcsNotificationOK44":
			err = notifications.ParseNotification(x.Data, db)
			if err != nil {
				atomic.AddInt32(&counters.withErrors, 1)
				atomic.AddInt32(&counters.notificationOKwithErrors, 1)
			} else {
				atomic.AddInt32(&counters.notifications, 1)
				atomic.AddInt32(&counters.notificationOK, 1)
			}
			break
		/*case "fcsNotificationPOT44":
			err = notifications.ParseNotificationOK(x.Data, db)
			if err != nil {
				atomic.AddInt32(&counters.withErrors, 1)
				atomic.AddInt32(&counters.notificationPOTwithErrors, 1)
			} else {
				atomic.AddInt32(&counters.notifications, 1)
				atomic.AddInt32(&counters.notificationPOT, 1)
			}
			break //*/
		case "fcsNotificationZK44":
			err = notifications.ParseNotification(x.Data, db)
			if err != nil {
				atomic.AddInt32(&counters.withErrors, 1)
				atomic.AddInt32(&counters.notificationZKwithErrors, 1)
			} else {
				atomic.AddInt32(&counters.notifications, 1)
				atomic.AddInt32(&counters.notificationZK, 1)
			}
			break
		case "fcsNotificationZP44":
			err = notifications.ParseNotification(x.Data, db)
			if err != nil {
				atomic.AddInt32(&counters.withErrors, 1)
				atomic.AddInt32(&counters.notificationZPwithErrors, 1)
			} else {
				atomic.AddInt32(&counters.notifications, 1)
				atomic.AddInt32(&counters.notificationZP, 1)
			}
			break
		case "fcsNotificationPO44":
			err = notifications.ParseNotification(x.Data, db)
			if err != nil {
				atomic.AddInt32(&counters.withErrors, 1)
				atomic.AddInt32(&counters.notificationPOwithErrors, 1)
			} else {
				atomic.AddInt32(&counters.notifications, 1)
				atomic.AddInt32(&counters.notificationPO, 1)
			}
			break
		case "fcsNotificationZKB44":
			err = notifications.ParseNotification(x.Data, db)
			if err != nil {
				atomic.AddInt32(&counters.withErrors, 1)
				atomic.AddInt32(&counters.notificationZKBwithErrors, 1)
			} else {
				atomic.AddInt32(&counters.notifications, 1)
				atomic.AddInt32(&counters.notificationZKB, 1)
			}
			break
		case "fcsNotificationISM44":
			err = notifications.ParseNotification(x.Data, db)
			if err != nil {
				atomic.AddInt32(&counters.withErrors, 1)
				atomic.AddInt32(&counters.notificationISMwithErrors, 1)
			} else {
				atomic.AddInt32(&counters.notifications, 1)
				atomic.AddInt32(&counters.notificationISM, 1)
			}
			break
		case "fcsNotificationINM111":
			err = notifications.ParseNotification(x.Data, db)
			if err != nil {
				atomic.AddInt32(&counters.withErrors, 1)
				atomic.AddInt32(&counters.notificationINM111withErrors, 1)
			} else {
				atomic.AddInt32(&counters.notifications, 1)
				atomic.AddInt32(&counters.notificationINM111, 1)
			}
			break
		case "fcsNotificationZKKU44":
			err = notifications.ParseNotification(x.Data, db)
			if err != nil {
				atomic.AddInt32(&counters.withErrors, 1)
				atomic.AddInt32(&counters.notificationZKKUwithErrors, 1)
			} else {
				atomic.AddInt32(&counters.notifications, 1)
				atomic.AddInt32(&counters.notificationZKKU, 1)
			}
			break
		case "fcsNotificationZKK44":
			err = notifications.ParseNotification(x.Data, db)
			if err != nil {
				atomic.AddInt32(&counters.withErrors, 1)
				atomic.AddInt32(&counters.notificationZKKwithErrors, 1)
			} else {
				atomic.AddInt32(&counters.notifications, 1)
				atomic.AddInt32(&counters.notificationZKK, 1)
			}
			break
		case "fcsNotificationZA44":
			err = notifications.ParseNotification(x.Data, db)
			if err != nil {
				atomic.AddInt32(&counters.withErrors, 1)
				atomic.AddInt32(&counters.notificationZAwithErrors, 1)
			} else {
				atomic.AddInt32(&counters.notifications, 1)
				atomic.AddInt32(&counters.notificationZA, 1)
			}
			break
		case "fcsPlacementResult":
			err = notifications.ParsePlacementResult(x.Data, db)
			if err != nil {
				atomic.AddInt32(&counters.withErrors, 1)
				atomic.AddInt32(&counters.notificationPlacementwithErrors, 1)
			} else {
				atomic.AddInt32(&counters.notificationPlacement, 1)
			}
			break
		/*
		case "fcsNotificationZakKD44":
			err = notifications.ParseNotificationZakKD(x.Data, db)
			break
		//*/
		case "fcsContractSign":
			/*err := ParseContractSign(x.Data, db)
			if err != nil {
				log.Warnln("ERROR: ",x.Name,err)
			}*/
			ignore("ContractSign:" + x.Name)
			break
		default:
			atomic.AddInt32(&counters.unsupported, 1)
			//log.Info("Force0 pool")
			err = errors.New("Unsupported xml type! " + x.Type)
			break
			//<-pool
			//sc.Done()
			//return nil
	}

	if err == nil {
		atomic.AddInt32(&counters.total, 1)
		//<-pool
		//sc.Done()
		if flagIsSave == 1 {
			err = SaveXml(x, db)
		}
		//return err
	} else {
		if strings.Split(err.Error(), " ")[0] != "Unsupported" {
			log.Warnln("ERROR: ", string(x.Bucket()), x.Name, err)
		}

		my_query := fmt.Sprint(`
			INSERT INTO parse_error_log(error_text, tag, zip_name, xml_name)
			VALUES ('`, err.Error(), `','`, string(x.Bucket()), `','`, x.Zip, `','`, x.Name, `')
		`)

		_, err := db.Exec(my_query)
		if err != nil {
			log.Println("ERROR Insert error:", my_query)
		}
	}

	//debug.FreeOSMemory()

	atomic.AddInt32(&counters.contracts, -counters.contracts223)
	//log.Info("Force2 pool")

	log.Println("ParserWorker::End.:", count)
	return nil
}
