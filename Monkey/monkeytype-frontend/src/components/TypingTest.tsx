import React, { useState, useEffect } from 'react';
import { sendTypingData } from '../services/api';

const TypingTest: React.FC = () => {
    const [textToType, setTextToType] = useState<string>('Type this text as fast as you can!');
    const [userInput, setUserInput] = useState<string>('');
    const [isTyping, setIsTyping] = useState<boolean>(false);
    const [startTime, setStartTime] = useState<number | null>(null);
    const [elapsedTime, setElapsedTime] = useState<number>(0);

    useEffect(() => {
        if (isTyping && startTime) {
            const interval = setInterval(() => {
                setElapsedTime(Math.floor((Date.now() - startTime) / 1000));
            }, 1000);
            return () => clearInterval(interval);
        }
    }, [isTyping, startTime]);

    const handleInputChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        const value = event.target.value;
        setUserInput(value);

        if (!isTyping) {
            setIsTyping(true);
            setStartTime(Date.now());
        }

        if (value === textToType) {
            handleTestComplete();
        }
    };

    const handleTestComplete = async () => {
        setIsTyping(false);
        await sendTypingData({ time: elapsedTime });
        alert(`Test completed in ${elapsedTime} seconds!`);
        resetTest();
    };

    const resetTest = () => {
        setUserInput('');
        setElapsedTime(0);
        setStartTime(null);
    };

    return (
        <div>
            <h1>Typing Test</h1>
            <p>{textToType}</p>
            <input
                type="text"
                value={userInput}
                onChange={handleInputChange}
                disabled={isTyping && userInput !== textToType}
            />
            <p>{isTyping ? `Time: ${elapsedTime}s` : 'Start typing to begin!'}</p>
        </div>
    );
};

export default TypingTest;