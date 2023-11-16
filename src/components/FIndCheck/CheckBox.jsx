import { Button, Input } from 'antd';
import React, { useEffect, useState } from 'react';

function CheckBox() {

    const [allchecked, setAllChecked] = useState([]);
    const [isChecked, setIsChecked] = useState(false);

    // function handleChange(e) {
    //     setChecked(e.target.checked);
    // }
    // useEffect(() => {
    //     const storedState = localStorage.getItem('checkboxState');
    //     if (storedState) {
    //         setAllChecked(JSON.parse(storedState));
    //     }
    // }, []);

    function handleChangeAll(e) {
        if (e.target.checked) {
            setAllChecked([...allchecked, e.target.value]);
        } else {
            setAllChecked(allchecked.filter((item) => item !== e.target.value));
        }
        setIsChecked(true);
        console.log("All Checked:", allchecked);
    }

    function handleSave() {
        if (isChecked) {
            const uniqueChecked = Array.from(new Set(allchecked)); // Remove duplicates
            localStorage.setItem('checkboxState', JSON.stringify(uniqueChecked));
            alert('State saved to localStorage!');
        } else {
            alert('No checkbox selected to save!');
        }
    }

    function handleLoad() {
        // Load state from localStorage
        const storedState = localStorage.getItem('checkboxState');
        if (storedState) {
            const loadedState = JSON.parse(storedState);
            setAllChecked(loadedState);
            setIsChecked(true);
            // Check checkboxes based on the loaded state
            document.querySelectorAll('input[type="checkbox"]').forEach((checkbox) => {
                checkbox.checked = loadedState.includes(checkbox.value);
            });
            alert('State loaded from localStorage!');
        } else {
            alert('No saved state found in localStorage.');
        }
    }


    console.log(localStorage.getItem('checkboxState'));
    return (
        <div className="form-check">
            <div>bao gồm các dịch vụ </div>
            <div>
                <input value="1" type="checkbox" onChange={handleChangeAll} checked={allchecked.includes("One")} />
                <span> dịch vụ  1 </span>
            </div>
            <div>
                <input value="2" type="checkbox" onChange={handleChangeAll} checked={allchecked.includes("Two")} />
                <span> dịch vụ 2 </span>
            </div>
            <div>
                <input value="3" type="checkbox" onChange={handleChangeAll} checked={allchecked.includes("Three")} />
                <span> dịch vụ  3 </span>
            </div>

            {/* <div>The all checked values are {allchecked.join(" , ")}</div> */}
            <Button type="primary" onClick={handleSave} disabled={allchecked.length === 0}>
                Save State
            </Button>
            <Button type="default" onClick={handleLoad} style={{ marginLeft: '10px' }}>
                Load State
            </Button>
            <h3>Display values</h3>
            <div>{allchecked}</div>
            <div>Return Json{}</div>
           

        </div>
    );
}

export default CheckBox;