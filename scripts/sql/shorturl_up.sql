-- @BeeOverwrite YES
-- @BeeGenerateTime 20201012_230803
CREATE TABLE `shorturl` (
     
        
            `id` int(11) NOT NULL AUTO_INCREMENT,
        
     
        
            `title` varchar(255) NOT NULL,
        
     
        
            `summary_id` int(11) NOT NULL,
        
     
        
            `url` varchar(255) NOT NULL,
        
     
     PRIMARY KEY (`id`)
)
