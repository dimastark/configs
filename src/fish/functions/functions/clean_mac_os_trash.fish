function clean_mac_os_trash
	cd /Applications

	sudo rm -rf Automator.app \
	            Calendar.app \
	            Chess.app \
	            Contacts.app \
	            Dashboard.app \
	            Dictionary.app \
	            "DVD Player.app" \
	            FaceTime.app \
	            "Image Capture.app" \
	            Maps.app \
	            Messages.app \
	            News.app \
	            Notes.app \
	            "Photo Booth.app" \
	            Photos.app \
	            Preview.app \
	            "QuickTime Player.app" \
	            Reminders.app \
	            Safari.app \
	            Siri.app \
	            Stickies.app \
	            Stocks.app \
	            VoiceMemos.app

	cd /Utilities

	sudo rm -rf "Boot Camp Assistant.app" \
	            "Feedback Assistant.app" \
	            "Grapher.app" \
	            "Migration Assistant.app" \
	            "Script Editor.app" \
	            "VoiceOver Utility.app"
end
