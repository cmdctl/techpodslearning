import React, {useEffect} from 'react';

export default function HomePage({
    load = () => {},
}) {
    useEffect(() => {
        load();
    }, [load])
    return (
        <div>this is home</div>
    )
}