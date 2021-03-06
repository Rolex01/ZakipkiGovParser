package main

// Stats contains atomic counters of processed documents
type Stats struct {
	contracts                 			int32
	contractswithErrors       			int32
	contractsProc             			int32
	contractsProcwithErrors   			int32
	contractsCancel           			int32
	contractsInProcess        			int32
	contractsProcInProcess    			int32
	contractsCancelInProcess  			int32

	notifications             			int32
	notificationEF            			int32
	notificationOK            			int32
	notificationEP            			int32
	notificationOKD           			int32
	notificationOKOU          			int32
	notificationOKU           			int32
	notificationPOT           			int32
	notificationPO            			int32
	notificationZK            			int32
	notificationZKB           			int32
	notificationZP            			int32
	notificationISM           			int32
	notificationINM111					int32
	notificationZKKU					int32
	notificationZKK						int32
	notificationZA						int32
	notificationPlacement				int32
	notificationEFwithErrors  			int32
	notificationOKwithErrors  			int32
	notificationEPwithErrors  			int32
	notificationOKDwithErrors 			int32
	notificationOKOUwithErrors			int32
	notificationOKUwithErrors 			int32
	notificationPOTwithErrors 			int32
	notificationPOwithErrors  			int32
	notificationZKwithErrors  			int32
	notificationZKBwithErrors 			int32
	notificationZPwithErrors  			int32
	notificationISMwithErrors 			int32
	notificationINM111withErrors		int32
	notificationZKKUwithErrors			int32
	notificationZKKwithErrors			int32
	notificationZAwithErrors			int32
	notificationPlacementwithErrors		int32

	contracts223              			int32
	unsupported               			int32
	total                     			int32
	withErrors                			int32

	nsi									int32
	nsiOrganizationType					int32
	nsiOrganization						int32
	nsiBudget							int32
	nsiOKEI								int32
	nsiOKPD								int32
	nsiOKPD2							int32
	nsiOKTMO							int32
	nsiOffBudget						int32
	nsiOffBudgetType					int32
	nsiPurchasePreference				int32
	nsiETPs								int32
	nsiPlacingWayList					int32
	nsiFarmDrugsDictionary				int32
	nsiOrganizationTypewithErrors		int32
	nsiOrganizationwithErrors			int32
	nsiBudgetwithErrors					int32
	nsiOKEIwithErrors					int32
	nsiOKPDwithErrors					int32
	nsiOKPD2withErrors					int32
	nsiOKTMOwithErrors					int32
	nsiOffBudgetwithErrors				int32
	nsiOffBudgetTypewithErrors			int32
	nsiPurchasePreferencewithErrors		int32
	nsiETPswithErrors					int32
	nsiPlacingWayListwithErrors			int32
	nsiFarmDrugsDictionarywithErrors	int32
}
