export namespace main {
	
	export class UserPermissionResult {
	    Name: string;
	    PermissionName: string;
	    ObjectName?: string;
	
	    static createFrom(source: any = {}) {
	        return new UserPermissionResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Name = source["Name"];
	        this.PermissionName = source["PermissionName"];
	        this.ObjectName = source["ObjectName"];
	    }
	}
	export class QueryResult[main.UserPermissionResult] {
	    Duration: number;
	    Data: UserPermissionResult[];
	
	    static createFrom(source: any = {}) {
	        return new QueryResult[main.UserPermissionResult](source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Duration = source["Duration"];
	        this.Data = this.convertValues(source["Data"], UserPermissionResult);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
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

