class SiteCache {
    cachedData: any;
    lastUsed: Date;
    lifeTime: number;
    key: string;

    constructor() {
        this.cachedData = null;
        this.lastUsed = new Date();
        this.lifeTime = 1000 * 30;
        this.key = "";
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

    set(key: string, data: any) {
        this.cachedData = data;
        this.lastUsed = new Date();
        this.key = key;
    }
}

export default class Cache {
    cache: SiteCache[];

    constructor() {
        this.cache = [];
    }

    get(key: string) {
        for (let i = 0; i < this.cache.length; i++) {
            if (this.cache[i].key === key) {
                return this.cache[i].get();
            }
        }
        return null;
    }

    set(key: string, data: any) {
        for (let i = 0; i < this.cache.length; i++) {
            if (this.cache[i].key === key) {
                this.cache[i].set(key, data);
                return;
            }
        }
        let siteCache = new SiteCache();
        siteCache.set(key, data);
        this.cache.push(siteCache);
    }
}
