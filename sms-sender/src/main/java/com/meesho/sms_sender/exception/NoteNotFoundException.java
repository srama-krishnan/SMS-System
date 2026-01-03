package com.meesho.sms_sender.exception;

public class NoteNotFoundException extends RuntimeException {
    public NoteNotFoundException(String id) {
        super("Note not found: " + id);
    }
}
