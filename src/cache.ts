export default class Cache {
    cachedData: any;
    lastUsed: Date;
    lifeTime: number;
    mensa: string;

    constructor() {
        this.cachedData = null;
        this.lastUsed = new Date();
        this.lifeTime = 1000 * 30;
        this.mensa = "";
    }

    get() {
        if (Date.now() - this.lastUsed.getTime() > this.lifeTime) {
            return null;
        }

        if (this.cachedData == null) {
            return null;
        }

        return this.cachedData;
    }

    set(mensa: string, data: any) {
        this.cachedData = data;
        this.lastUsed = new Date();
        this.mensa = mensa;
    }
}

