function clean_mac_os_trash
    set P $PWD

    cd /Applications

    sudo rm -rf Stocks.app \
            VoiceMemos.app \
            Home.app \
            Notes.app \
            Stickies.app \
            "Image Capture.app" \
            Maps.app \
            Contacts.app \
            Reminders.app \
            Dictionary.app \
            Messages.app \
            Photos.app \
            Chess.app \
            Automator.app \
            Dashboard.app \
            FaceTime.app \
            iTunes.app \
            "Mission Control.app" \
            "Photo Booth.app" \
            Safari.app \
            Siri.app \
            "Time Machine.app"

    cd /Applications/Utilities

    sudo rm -rf "Audio MIDI Setup.app" \
            "Boot Camp Assistant.app" \
            Console.app \
            "ColorSync Utility.app" \
            "Feedback Assistant.app" \
            "Grapher.app" \
            "Migration Assistant.app" \
            "Script Editor.app" \
            "VoiceOver Utility.app"
    cd $P	
end
