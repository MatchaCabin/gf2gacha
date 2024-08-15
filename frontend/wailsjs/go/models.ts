export namespace model {
	
	export class CommunityExchangeList {
	    id: number;
	    name: string;
	
	    static createFrom(source: any = {}) {
	        return new CommunityExchangeList(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	    }
	}
	export class DisplayRecord {
	    id: number;
	    name: string;
	    lose: boolean;
	    count: number;
	
	    static createFrom(source: any = {}) {
	        return new DisplayRecord(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.lose = source["lose"];
	        this.count = source["count"];
	    }
	}
	export class LogInfo {
	    tablePath: string;
	    accessToken: string;
	    uid: string;
	    gachaUrl: string;
	
	    static createFrom(source: any = {}) {
	        return new LogInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.tablePath = source["tablePath"];
	        this.accessToken = source["accessToken"];
	        this.uid = source["uid"];
	        this.gachaUrl = source["gachaUrl"];
	    }
	}
	export class Pool {
	    poolType: number;
	    gachaCount: number;
	    loseCount: number;
	    guaranteesCount: number;
	    rank5Count: number;
	    rank4Count: number;
	    rank3Count: number;
	    storedCount: number;
	    recordList: DisplayRecord[];
	
	    static createFrom(source: any = {}) {
	        return new Pool(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.poolType = source["poolType"];
	        this.gachaCount = source["gachaCount"];
	        this.loseCount = source["loseCount"];
	        this.guaranteesCount = source["guaranteesCount"];
	        this.rank5Count = source["rank5Count"];
	        this.rank4Count = source["rank4Count"];
	        this.rank3Count = source["rank3Count"];
	        this.storedCount = source["storedCount"];
	        this.recordList = this.convertValues(source["recordList"], DisplayRecord);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

