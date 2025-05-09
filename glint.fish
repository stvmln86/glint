# # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # #
#          glint · minimal note-taking functions in fish · by Stephen Malone          #
# # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # #

# # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # #
#                      part one · environment variable functions                      #
# # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # #

function _g_evar_dire -d "Echo a resolved GLINT_DIR."
	echo (path resolve $GLINT_DIR)
end

function _g_evar_extn -d "Echo a lowercased no-dot GLINT_EXT."
	set trim (string trim --left --chars "." $GLINT_EXT)
	echo (string lower $trim)
end

# # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # #
#                     part two · file path manipulation functions                     #
# # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # #

function _g_path_join -d "Echo a joined dire/name.extn path." -a dire name extn
	echo "$dire/$name.$extn"
end

function _g_path_name -d "Echo a lowercased path name." -a path
	set name (path basename $path --no-extension)
	echo (string lower $name)
end

# # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # #
#                     part three · file system handling functions                     #
# # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # #

# # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # #
#                           part four · main glint functions                          #
# # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # #

function gls -d "List all notes, or notes starting with TEXT." -a text
	set dire (_g_evar_dire)
	set extn (_g_evar_extn)
	
	for path in $dire/*.$extn
		set name (_g_path_name $path)
		set okay (string match --entire "$text*" $name)

		if test $okay = $name
			echo $name
		end
	end
end

function gop -d "Open NOTE in $VISUAL." -a note
	set dire (_g_evar_dire)
	set extn (_g_evar_extn)
	set name (_g_path_name $note) 
	set path (_g_path_join $dire $name $extn)
	eval "$VISUAL $path"
end
	
