import axios from 'axios';

const isUrl = (url: string) => {
    try {
        new URL(url);
        return true;
    } catch (err) {
        return false;
    }
}

export default async function fetch(url: string) {
    if (!isUrl(url)) {
        throw new Error('Invalid URL');
    }
    try {
        const res = await axios.get(url)
        return res.data;
    } catch (err) {
        console.log(err);
        return null;
    }
}
