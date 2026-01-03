package com.meesho.sms_sender.service;

import com.meesho.sms_sender.dto.CreateNoteRequest;
import com.meesho.sms_sender.dto.Note;
import com.meesho.sms_sender.exception.NoteNotFoundException;
import org.springframework.stereotype.Service;

import java.time.Instant;
import java.util.*;
import java.util.concurrent.ConcurrentHashMap;

@Service
public class NoteService {

    private final Map<String, Note> notes = new ConcurrentHashMap<>();

    public Note create(CreateNoteRequest req) {
        String id = UUID.randomUUID().toString();
        Note note = new Note(id, req.getTitle(), req.getContent(), Instant.now());
        notes.put(id, note);
        return note;
    }

    public Note getById(String id) {
        Note note = notes.get(id);
        if (note == null) throw new NoteNotFoundException(id);
        return note;
    }

    public List<Note> getAll() {
        return new ArrayList<>(notes.values());
    }

    public void delete(String id) {
        if (notes.remove(id) == null) throw new NoteNotFoundException(id);
    }
}
