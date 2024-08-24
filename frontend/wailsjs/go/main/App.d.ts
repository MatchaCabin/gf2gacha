// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {model} from '../models';

export function ApplyUpdate():Promise<string>;

export function CheckUpdate():Promise<string>;

export function ExportRawJson(arg1:string):Promise<string>;

export function GetCommunityExchangeList():Promise<Array<model.CommunityExchangeList>>;

export function GetLogInfo():Promise<model.LogInfo>;

export function GetPoolInfo(arg1:string,arg2:number):Promise<model.Pool>;

export function GetSettingExchangeList():Promise<Array<number>>;

export function GetUserList():Promise<Array<string>>;

export function HandleCommunityTasks():Promise<Array<string>>;

export function ImportRawJson(arg1:string,arg2:boolean):Promise<string>;

export function MergeEreRecord(arg1:string,arg2:string):Promise<string>;

export function SaveSettingExchangeList(arg1:Array<number>):Promise<void>;

export function UpdatePoolInfo(arg1:boolean):Promise<Array<string>>;
