<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: v1/subscribe/subscribe.proto

namespace GPBMetadata\V1\Subscribe;

class Subscribe
{
    public static $is_initialized = false;

    public static function initOnce() {
        $pool = \Google\Protobuf\Internal\DescriptorPool::getGeneratedPool();

        if (static::$is_initialized == true) {
          return;
        }
        \GPBMetadata\Google\Api\Annotations::initOnce();
        \GPBMetadata\V1\Payment\Payment::initOnce();
        $pool->internalAddGeneratedFile(hex2bin(
            "0aa71d0a1c76312f7375627363726962652f7375627363726962652e70726f746f12097375627363726962651a1876312f7061796d656e742f7061796d656e742e70726f746f228f030a154f6e6574696d655061796d656e745265717565737412140a0c6d65726368616e745f756964180120012809120e0a06616d6f756e7418022001280512100a087461785f6672656518032001280512130a0b636172645f6e756d626572180420012809120e0a06657870697279180520012809120d0a05626972746818062001280912120a0a7077645f32646967697418072001280912140a0c637573746f6d65725f756964180820012809120a0a027067180920012809120c0a046e616d65180a2001280912120a0a62757965725f6e616d65180b2001280912130a0b62757965725f656d61696c180c2001280912110a0962757965725f74656c180d2001280912120a0a62757965725f61646472180e2001280912160a0e62757965725f706f7374636f6465180f2001280912120a0a636172645f71756f746118102001280512210a19696e7465726573745f667265655f62795f6d65726368616e7418112001280812130a0b637573746f6d5f6461746118122001280912120a0a6e6f746963655f75726c181320012809225b0a164f6e6574696d655061796d656e74526573706f6e7365120c0a04636f6465180120012805120f0a076d65737361676518022001280912220a08726573706f6e736518032001280b32102e7061796d656e742e5061796d656e7422b9020a13416761696e5061796d656e745265717565737412140a0c637573746f6d65725f75696418012001280912140a0c6d65726368616e745f756964180220012809120e0a06616d6f756e7418032001280512100a087461785f66726565180420012805120c0a046e616d6518052001280912120a0a62757965725f6e616d6518062001280912130a0b62757965725f656d61696c18072001280912110a0962757965725f74656c18082001280912120a0a62757965725f6164647218092001280912160a0e62757965725f706f7374636f6465180a2001280912120a0a636172645f71756f7461180b2001280512210a19696e7465726573745f667265655f62795f6d65726368616e74180c2001280812130a0b637573746f6d5f64617461180d2001280912120a0a6e6f746963655f75726c180e2001280922590a14416761696e5061796d656e74526573706f6e7365120c0a04636f6465180120012805120f0a076d65737361676518022001280912220a08726573706f6e736518032001280b32102e7061796d656e742e5061796d656e7422d9010a145061796d656e745363686564756c65506172616d12140a0c6d65726368616e745f75696418012001280912130a0b7363686564756c655f6174180220012805120e0a06616d6f756e7418032001280512100a087461785f66726565180420012805120c0a046e616d6518052001280912120a0a62757965725f6e616d6518062001280912130a0b62757965725f656d61696c18072001280912110a0962757965725f74656c18082001280912120a0a62757965725f6164647218092001280912160a0e62757965725f706f7374636f6465180a2001280922f9020a1b556e69745363686564756c655061796d656e74526573706f6e736512140a0c637573746f6d65725f75696418012001280912140a0c6d65726368616e745f756964180220012809120f0a07696d705f75696418032001280912130a0b7363686564756c655f617418042001280512130a0b65786563757465645f617418052001280512120a0a7265766f6b65645f6174180620012805120e0a06616d6f756e74180720012805120c0a046e616d6518082001280912120a0a62757965725f6e616d6518092001280912130a0b62757965725f656d61696c180a2001280912110a0962757965725f74656c180b2001280912120a0a62757965725f61646472180c2001280912160a0e62757965725f706f7374636f6465180d2001280912130a0b637573746f6d5f64617461180e2001280912170a0f7363686564756c655f737461747573180f2001280912160a0e7061796d656e745f73746174757318102001280912130a0b6661696c5f726561736f6e18112001280922cf010a165363686564756c65506179656d6e745265717565737412140a0c637573746f6d65725f75696418012001280912170a0f636865636b696e675f616d6f756e7418022001280512130a0b636172645f6e756d626572180320012809120e0a06657870697279180420012809120d0a05626972746818052001280912120a0a7077645f326469676974180620012809120a0a02706718072001280912320a097363686564756c657318082003280b321f2e7375627363726962652e5061796d656e745363686564756c65506172616d22720a175363686564756c655061796d656e74526573706f6e7365120c0a04636f6465180120012805120f0a076d65737361676518022001280912380a08726573706f6e736518032003280b32262e7375627363726962652e556e69745363686564756c655061796d656e74526573706f6e736522460a18556e7363686564756c655061796d656e745265717565737412140a0c637573746f6d65725f75696418012001280912140a0c6d65726368616e745f75696418022003280922740a19556e7363686564756c655061796d656e74526573706f6e7365120c0a04636f6465180120012805120f0a076d65737361676518022001280912380a08726573706f6e736518032003280b32262e7375627363726962652e556e69745363686564756c655061796d656e74526573706f6e736522310a194765745061796d656e745363686564756c655265717565737412140a0c6d65726368616e745f75696418012001280922750a1a4765745061796d656e745363686564756c65526573706f6e7365120c0a04636f6465180120012805120f0a076d65737361676518022001280912380a08726573706f6e736518032001280b32262e7375627363726962652e556e69745363686564756c655061796d656e74526573706f6e7365227c0a234765745061796d656e745363686564756c654279437573746f6d65725265717565737412140a0c637573746f6d65725f756964180120012809120c0a0470616765180220012805120c0a0466726f6d180320012805120a0a02746f18042001280512170a0f7363686564756c655f737461747573180520012809228d010a264e65737465644765745061796d656e745363686564756c654279437573746f6d657244617461120d0a05746f74616c18012001280512100a0870726576696f7573180220012805120c0a046e65787418032001280512340a046c69737418042003280b32262e7375627363726962652e556e69745363686564756c655061796d656e74526573706f6e7365228a010a244765745061796d656e745363686564756c654279437573746f6d6572526573706f6e7365120c0a04636f6465180120012805120f0a076d65737361676518022001280912430a08726573706f6e736518032001280b32312e7375627363726962652e4e65737465644765745061796d656e745363686564756c654279437573746f6d65724461746132eb070a10537562736372696265536572766963651290010a114f6e6574696d655061796d656e7452504312202e7375627363726962652e4f6e6574696d655061796d656e74526571756573741a212e7375627363726962652e4f6e6574696d655061796d656e74526573706f6e7365223682d3e4930230222b2f6170692f7061796d656e74732f76312f7375627363726962652f7061796d656e74732f6f6e6574696d653a012a1288010a0f416761696e5061796d656e74525043121e2e7375627363726962652e416761696e5061796d656e74526571756573741a1f2e7375627363726962652e416761696e5061796d656e74526573706f6e7365223482d3e493022e22292f6170692f7061796d656e74732f76312f7375627363726962652f7061796d656e74732f616761696e3a012a1294010a125363686564756c655061796d656e7452504312212e7375627363726962652e5363686564756c65506179656d6e74526571756573741a222e7375627363726962652e5363686564756c655061796d656e74526573706f6e7365223782d3e4930231222c2f6170692f7061796d656e74732f76312f7375627363726962652f7061796d656e74732f7363686564756c653a012a129c010a14556e7363686564756c655061796d656e7452504312232e7375627363726962652e556e7363686564756c655061796d656e74526571756573741a242e7375627363726962652e556e7363686564756c655061796d656e74526573706f6e7365223982d3e4930233222e2f6170692f7061796d656e74732f76312f7375627363726962652f7061796d656e74732f756e7363686564756c653a012a12aa010a164765745363686564756c65645061796d656e7452504312242e7375627363726962652e4765745061796d656e745363686564756c65526571756573741a252e7375627363726962652e4765745061796d656e745363686564756c65526573706f6e7365224382d3e493023d123b2f6170692f7061796d656e74732f76312f7375627363726962652f7061796d656e74732f7363686564756c652f7b6d65726368616e745f7569647d12d5010a234765745363686564756c65645061796d656e744279437573746f6d6572556964525043122e2e7375627363726962652e4765745061796d656e745363686564756c654279437573746f6d6572526571756573741a2f2e7375627363726962652e4765745061796d656e745363686564756c654279437573746f6d6572526573706f6e7365224d82d3e493024712452f6170692f7061796d656e74732f76312f7375627363726962652f7061796d656e74732f7363686564756c652f637573746f6d6572732f7b637573746f6d65725f7569647d42455a346769746875622e636f6d2f69616d706f72742f696e746572666163652f67656e5f7372632f676f2f76312f737562736372696265aa020c56312e537562736372696265620670726f746f33"
        ), true);

        static::$is_initialized = true;
    }
}

