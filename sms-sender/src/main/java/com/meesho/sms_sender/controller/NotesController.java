package com.meesho.sms_sender.controller;

import com.meesho.sms_sender.dto.CreateNoteRequest;
import com.meesho.sms_sender.dto.Note;
import com.meesho.sms_sender.service.NoteService;
import jakarta.validation.Valid;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@RequestMapping("/v1/notes")
public class NotesController {

    private final NoteService noteService;

    public NotesController(NoteService noteService) {
        this.noteService = noteService;
    }

    @PostMapping
    public ResponseEntity<Note> create(@Valid @RequestBody CreateNoteRequest req) {
        return ResponseEntity.ok(noteService.create(req));
    }

    @GetMapping("/{id}")
    public ResponseEntity<Note> get(@PathVariable String id) {
        return ResponseEntity.ok(noteService.getById(id));
    }

    @GetMapping
    public ResponseEntity<List<Note>> list() {
        return ResponseEntity.ok(noteService.getAll());
    }

    @DeleteMapping("/{id}")
    public ResponseEntity<Void> delete(@PathVariable String id) {
        noteService.delete(id);
        return ResponseEntity.noContent().build();
    }
}
