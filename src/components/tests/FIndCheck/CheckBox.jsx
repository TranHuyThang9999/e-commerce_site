import { Button, Input } from 'antd';
import React, { useEffect, useState } from 'react';

function CheckBox() {

    const [allchecked, setAllChecked] = useState([]);
    const [isChecked, setIsChecked] = useState(false);

    function handleChangeAll(e) {
        if (e.target.checked) {
            setAllChecked([...allchecked, e.target.value]);
        } else {
            setAllChecked(allchecked.filter((item) => item !== e.target.value));
        }
        setIsChecked(true);
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
    const jsonResult = {
        dichvu1: allchecked.includes('1') ? 1 : 0,
        dichvu2: allchecked.includes('2') ? 2 : 0,
        dichvu3: allchecked.includes('3') ? 3 : 0,
    };
    function handleReturnJson() {

        console.log('Return Json:', jsonResult);
    }

    return (
        <div className="form-check">
            <div>bao gồm các dịch vụ </div>
            <div>
                <input value="1" type="checkbox" onChange={handleChangeAll} checked={allchecked.includes("1")} />
                <span> dịch vụ  1 </span>
            </div>
            <div>
                <input value="2" type="checkbox" onChange={handleChangeAll} checked={allchecked.includes("2")} />
                <span> dịch vụ 2 </span>
            </div>
            <div>
                <input value="3" type="checkbox" onChange={handleChangeAll} checked={allchecked.includes("3")} />
                <span> dịch vụ  3 </span>
            </div>

            <Button type="primary" onClick={handleSave} disabled={allchecked.length === 0}>
                Save State
            </Button>
            <Button type="default" onClick={handleLoad} style={{ marginLeft: '10px' }}>
                Load State
            </Button>
            <Button type="default" onClick={handleReturnJson} style={{ marginLeft: '10px' }}>
                Return Json
            </Button>
            <p>Data</p>
            <pre>{JSON.stringify(jsonResult, null, 2)}</pre>
        </div>
    );
}

export default CheckBox;
