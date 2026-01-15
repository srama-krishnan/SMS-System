package com.meesho.sms_sender.dto;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

/**
 * DTO representing an SMS event to be sent to Kafka.
 * This structure matches the Go service's Message model for seamless integration.
 */
@Data
@NoArgsConstructor
@AllArgsConstructor
public class SmsEvent {
    /**
     * User ID who sent/received the SMS
     */
    private String userId;

    /**
     * Phone number (10 digits)
     */
    private String phoneNumber;

    /**
     * SMS message text content
     * Note: Named "text" to match Go's Message model field name
     */
    private String text;

    /**
     * SMS delivery status: "SUCCESS" or "FAIL"
     */
    private String status;

    /**
     * Timestamp when the SMS event occurred (milliseconds since epoch)
     * Note: Go's Message uses CreatedAt (time.Time), which will be converted from this timestamp
     */
    private long timestamp;
}
